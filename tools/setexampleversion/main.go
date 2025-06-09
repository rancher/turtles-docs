package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	version := flag.String("version", "", "version tag to replace with (e.g., v0.20.0)")
	flag.Parse()

	if *version == "" {
		fmt.Fprintf(os.Stderr, "Error: version flag is required\n")
		fmt.Fprintf(os.Stderr, "Usage: %s -version=v0.20.0 file1.adoc file2.adoc\n", os.Args[0])
		os.Exit(1)
	}

	files := flag.Args()
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "Error: no files specified\n")
		fmt.Fprintf(os.Stderr, "Usage: %s -version=v0.20.0 file1.adoc file2.adoc\n", os.Args[0])
		os.Exit(1)
	}

	if !strings.HasPrefix(*version, "v") {
		fmt.Fprintf(os.Stderr, "Error: version must start with 'v' (e.g., v0.20.0)\n")
		os.Exit(1)
	}

	totalReplacements := 0
	for _, file := range files {
		replacements := updateFile(file, *version)
		if replacements > 0 {
			fmt.Printf("Updated %d links in %s\n", replacements, file)
			totalReplacements += replacements
		}
	}

	fmt.Printf("Done. Total replacements: %d\n", totalReplacements)
}

func updateFile(filename, version string) int {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filename, err)
		return 0
	}

	originalContent := string(content)
	updatedContent := updateRawGitHubLinks(originalContent, version)

	if originalContent == updatedContent {
		return 0
	}

	err = os.WriteFile(filename, []byte(updatedContent), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file %s: %v\n", filename, err)
		return 0
	}

	return countReplacements(originalContent, updatedContent)
}

func updateRawGitHubLinks(content, version string) string {
	rawGitHubRegex := regexp.MustCompile(`https://raw\.githubusercontent\.com/rancher/turtles/refs/heads/main/`)

	replacement := fmt.Sprintf("https://raw.githubusercontent.com/rancher/turtles/refs/tags/%s/", version)

	return rawGitHubRegex.ReplaceAllString(content, replacement)
}

func countReplacements(original, updated string) int {
	originalLines := strings.Split(original, "\n")
	updatedLines := strings.Split(updated, "\n")

	replacements := 0
	for i := 0; i < len(originalLines) && i < len(updatedLines); i++ {
		if originalLines[i] != updatedLines[i] {
			replacements++
		}
	}

	return replacements
}
