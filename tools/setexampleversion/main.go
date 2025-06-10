package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type ReplacementRule struct {
	Name        string `json:"name"`
	Pattern     string `json:"pattern"`
	Replacement string `json:"replacement"`
	Version     string `json:"version,omitempty"`
}

type Config struct {
	DefaultVersion string            `json:"default_version,omitempty"`
	Rules          []ReplacementRule `json:"rules"`
}

func main() {
	version := flag.String("version", "", "default version to replace with (e.g., v0.20.0)")
	configFile := flag.String("config", "replace-rules.json", "JSON config file path")
	flag.Parse()

	configData, err := os.ReadFile(*configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config file %s: %v\n", *configFile, err)
		os.Exit(1)
	}

	var config Config
	if err := json.Unmarshal(configData, &config); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing config file: %v\n", err)
		os.Exit(1)
	}

	defaultVersion := *version
	if defaultVersion == "" && config.DefaultVersion != "" {
		defaultVersion = config.DefaultVersion
	}

	files := flag.Args()
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "Error: no files specified\n")
		fmt.Fprintf(os.Stderr, "Usage: %s -config=config.json -version=v0.20.0 file1.adoc file2.adoc\n", os.Args[0])
		os.Exit(1)
	}

	totalReplacements := 0
	for _, file := range files {
		fileReplacements := 0

		for _, rule := range config.Rules {
			ruleVersion := getRuleVersion(rule, defaultVersion)
			if ruleVersion == "" {
				fmt.Fprintf(os.Stderr, "Warning: No version specified for rule '%s' and no default version available\n", rule.Name)
				continue
			}
			replacements := updateFileWithRule(file, rule, ruleVersion)
			fileReplacements += replacements
		}

		if fileReplacements > 0 {
			fmt.Printf("Updated %d links in %s\n", fileReplacements, file)
			totalReplacements += fileReplacements
		}
	}

	fmt.Printf("Done. Total replacements: %d\n", totalReplacements)
}

func getRuleVersion(rule ReplacementRule, defaultVersion string) string {
	if rule.Version != "" {
		return rule.Version
	}
	return defaultVersion
}

func updateFileWithRule(filename string, rule ReplacementRule, version string) int {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filename, err)
		return 0
	}

	originalContent := string(content)
	updatedContent := applyRule(originalContent, rule, version)

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

func applyRule(content string, rule ReplacementRule, version string) string {
	regex, err := regexp.Compile(rule.Pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error compiling regex for rule %s: %v\n", rule.Name, err)
		return content
	}

	replacement := fmt.Sprintf(rule.Replacement, version)
	return regex.ReplaceAllString(content, replacement)
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
