package seed

import (
	"os"
	"path/filepath"
	"testing"
)

func createTmpFile(t *testing.T, content string) string {
	tmp := t.TempDir()
	file := filepath.Join(tmp, "urls.txt")
	err := os.WriteFile(file, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	return file
}

func TestReadURLs_ValidFile(t *testing.T) {
	content := `
https://example.com
https://golang.org
# this is a comment
https://news.ycombinator.com
	`
	file := createTmpFile(t, content)
	urls, err := ReadURLs(file)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := []string{
		"https://example.com",
		"https://golang.org",
		"https://news.ycombinator.com",
	}

	if len(urls) != len(expected) {
		t.Fatalf("expected %d URLs, got %d", len(expected), len(urls))
	}

	for i, url := range urls {
		if url != expected[i] {
			t.Errorf("expected %s, got %s", expected[i], url)
		}
	}
}

func TestReadURLs_EmptyFile(t *testing.T) {
	content := ""
	file := createTmpFile(t, content)
	urls, err := ReadURLs(file)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(urls) != 0 {
		t.Fatalf("expected 0 URLs, got %d", len(urls))
	}
}

func TestReadURLs_FileNotFound(t *testing.T) {
	file := "nonexistent.txt"
	urls, err := ReadURLs(file)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if urls != nil {
		t.Fatalf("expected nil URLs, got %v", urls)
	}
}
