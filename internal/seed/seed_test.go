package seed

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTmpFile(t *testing.T, content string) string {
	tmp := t.TempDir()
	file := filepath.Join(tmp, "urls.txt")
	err := os.WriteFile(file, []byte(content), 0644)
	assert.NoError(t, err, "failed to create temp file")
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
	assert.NoError(t, err)
	assert.Len(t, urls, 3)

	expected := []string{
		"https://example.com",
		"https://golang.org",
		"https://news.ycombinator.com",
	}

	assert.Equal(t, expected, urls)
}

func TestReadURLs_EmptyFile(t *testing.T) {
	file := createTmpFile(t, "")

	urls, err := ReadURLs(file)
	assert.NoError(t, err)
	assert.Empty(t, urls)
}

func TestReadURLs_FileNotFound(t *testing.T) {
	urls, err := ReadURLs("nonexistent.txt")
	assert.Error(t, err)
	assert.Nil(t, urls)
}
