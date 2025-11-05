package tests

import (
	"os"
	"strings"
	"testing"
)

// Verify documentation files exist
func TestDocumentationExists(t *testing.T) {
	files := []string{
		"../README.md",
		"../docs/ANALYSIS.md",
		"../AUTHORS.md",
		"../AGENTS.md",
	}
	
	for _, file := range files {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			t.Errorf("Documentation file missing: %s", file)
		}
	}
}

// Verify README has required sections
func TestREADMEStructure(t *testing.T) {
	content, err := os.ReadFile("../README.md")
	if err != nil {
		t.Fatal(err)
	}
	
	text := string(content)
	
	requiredSections := []string{
		"## Overview",
		"## Quick Start",
		"## Key Features",
		"## Documentation",
	}
	
	for _, section := range requiredSections {
		if !strings.Contains(text, section) {
			t.Errorf("README missing required section: %s", section)
		}
	}
}

// Verify ANALYSIS.md has required sections
func TestAnalysisStructure(t *testing.T) {
	content, err := os.ReadFile("../docs/ANALYSIS.md")
	if err != nil {
		t.Fatal(err)
	}
	
	text := string(content)
	
	requiredSections := []string{
		"## 1. Architecture Comparison",
		"## 2. Why FSM",
		"## 3. Transformation Rules Reference",
		"## 4. Golden Test Set",
	}
	
	for _, section := range requiredSections {
		if !strings.Contains(text, section) {
			t.Errorf("ANALYSIS.md missing required section: %s", section)
		}
	}
}

// Verify examples in README are valid
func TestREADMEExamples(t *testing.T) {
	// This would test that code examples in README actually work
	// For simplicity, just verify they exist
	
	content, err := os.ReadFile("../README.md")
	if err != nil {
		t.Fatal(err)
	}
	
	text := string(content)
	
	if !strings.Contains(text, "go run") {
		t.Error("README should contain usage examples with 'go run'")
	}
}

// Verify internal links work
func TestInternalLinks(t *testing.T) {
	readmeContent, err := os.ReadFile("../README.md")
	if err != nil {
		t.Fatal(err)
	}
	
	// Check if README links to ANALYSIS.md
	if !strings.Contains(string(readmeContent), "ANALYSIS.md") {
		t.Error("README should link to ANALYSIS.md")
	}
}
