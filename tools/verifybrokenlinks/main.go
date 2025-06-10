package main

import (
	"flag"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type URLOccurrence struct {
	URL  string
	File string
}

type URLCheckResult struct {
	URL        string
	IsValid    bool
	StatusCode int
	ErrorMsg   string
}

func main() {
	docsDir := flag.String("docs-dir", "docs", "directory to scan for .adoc files")
	maxParallel := flag.Int("max-parallel", 10, "maximum number of parallel link checks")
	exitOnError := flag.Bool("exit-on-error", true, "exit with non-zero code if broken links are found")
	flag.Parse()

	allOccurrences, uniqueURLs := extractLinks(*docsDir)
	results := checkLinks(uniqueURLs, *maxParallel, len(uniqueURLs))

	brokenLinks := make(map[string]URLCheckResult)
	for _, result := range results {
		if !result.IsValid {
			brokenLinks[result.URL] = result
		}
	}

	for url, result := range brokenLinks {
		errorCode := result.ErrorMsg
		if result.StatusCode != 0 {
			errorCode = fmt.Sprintf("HTTP %d", result.StatusCode)
		}
		fmt.Printf("❌ %s %s\nFound in:\n", errorCode, url)
		for _, occurrence := range allOccurrences {
			if occurrence.URL == url {
				fmt.Printf("- %s\n", strings.TrimPrefix(occurrence.File, *docsDir+"/"))
			}
		}
		fmt.Println()
	}

	rateLimited := make(map[string]URLCheckResult)
	actualBroken := make(map[string]URLCheckResult)
	for url, result := range brokenLinks {
		if result.StatusCode == 429 || result.ErrorMsg == "Rate limited (429)" {
			rateLimited[url] = result
		} else {
			actualBroken[url] = result
		}
	}

	fmt.Printf("Found %d broken links out of %d total links\n", len(actualBroken), len(uniqueURLs))
	if len(rateLimited) > 0 {
		fmt.Printf("Additionally, %d links were rate limited (429)\n", len(rateLimited))
	}
	fmt.Println("Done")

	if len(actualBroken) > 0 && *exitOnError {
		os.Exit(1)
	} else if len(brokenLinks) == 0 {
		fmt.Printf("Success: All %d links are valid\n", len(uniqueURLs))
	}
}

func extractLinks(docsDir string) ([]URLOccurrence, []string) {
	var occurrences []URLOccurrence
	uniqueURLs := make(map[string]struct{})

	githubLinkRegex := regexp.MustCompile(`https?://(?:[a-zA-Z0-9-]+\.)*github(?:usercontent)?\.com[-a-zA-Z0-9@:%._\+~#=/]*`)
	issuesPRRegex := regexp.MustCompile(`/(?:issues|pull)/[0-9]+`)
	componentsRegex := regexp.MustCompile(`/releases/(?:v|tag/v)?[0-9]+\.[0-9]+\.[0-9]+/.*-components\.ya?ml$`)

	filepath.WalkDir(docsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".adoc") {
			return err
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		for _, url := range githubLinkRegex.FindAllString(string(content), -1) {
			cleaned := trimTrailingPunctuation(url)
			if !issuesPRRegex.MatchString(cleaned) && !componentsRegex.MatchString(cleaned) { // Ignore issues, PRs, and component YAML links
				occurrences = append(occurrences, URLOccurrence{URL: cleaned, File: path})
				uniqueURLs[cleaned] = struct{}{}
			}
		}
		return nil
	})

	urls := make([]string, 0, len(uniqueURLs))
	for url := range uniqueURLs {
		urls = append(urls, url)
	}

	return occurrences, urls
}

func trimTrailingPunctuation(url string) string {
	return strings.TrimRight(url, ".,)]>")
}

func checkLinks(urls []string, maxParallel int, total int) []URLCheckResult {
	resultsChan := make(chan URLCheckResult, len(urls))
	sem := make(chan struct{}, maxParallel)
	var wg sync.WaitGroup
	var progress int32

	rateLimiter := make(chan struct{}, 1) // Without rate limiting, scripts get blocked by 429 error
	go func() {
		for {
			rateLimiter <- struct{}{}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	client := &http.Client{
		Timeout: 10 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
	}

	for _, url := range urls {
		wg.Add(1)
		sem <- struct{}{}
		go func(u string) {
			defer wg.Done()
			defer func() { <-sem }()
			resultsChan <- checkLinkWithRetry(client, rateLimiter, u)
			fmt.Fprintf(os.Stderr, "\rProgress: %d/%d links checked", atomic.AddInt32(&progress, 1), total)
		}(url)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	var results []URLCheckResult
	for result := range resultsChan {
		results = append(results, result)
	}

	fmt.Fprintln(os.Stderr)
	return results
}

func checkLinkWithRetry(client *http.Client, rateLimiter <-chan struct{}, url string) URLCheckResult {
	result := URLCheckResult{URL: url, IsValid: false}
	for attempt := 0; attempt < 3; attempt++ {
		if attempt > 0 {
			sleep := time.Duration(1<<uint(attempt)) * time.Second
			time.Sleep(sleep)
			fmt.Fprintf(os.Stderr, "\rRetrying %s (attempt %d/3) after %v", url, attempt+1, sleep)
		}
		<-rateLimiter

		req, err := http.NewRequest("HEAD", url, nil)
		if err != nil {
			result.ErrorMsg = err.Error()
			continue
		}

		resp, err := client.Do(req)
		if err != nil {
			<-rateLimiter
			resp, err = client.Get(url)
			if err != nil {
				result.ErrorMsg = "Connection failed or timed out"
				continue
			}
		}

		defer resp.Body.Close()
		result.StatusCode = resp.StatusCode

		if resp.StatusCode >= 200 && resp.StatusCode < 400 {
			result.IsValid = true
			break
		}
		if resp.StatusCode == 429 {
			if attempt < 2 {
				continue
			}
			result.ErrorMsg = "Rate limited (429)"
		} else {
			break
		}
	}
	return result
}
