package frontier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrontier_EnqueueDequeue(t *testing.T) {
	f := New()

	assert.True(t, f.IsEmpty(), "Frontier should be empty initially")

	f.Enqueue("https://example.com")
	f.Enqueue("https://golang.org")

	assert.False(t, f.IsEmpty(), "Frontier should not be empty after enqueue")

	url1 := f.Dequeue()
	url2 := f.Dequeue()

	assert.Equal(t, "https://example.com", url1)
	assert.Equal(t, "https://golang.org", url2)

	assert.True(t, f.IsEmpty(), "Frontier should be empty after dequeuing all elements")
}

func TestFrontier_Deduplication(t *testing.T) {
	f := New()

	f.Enqueue("https://example.com")
	f.Enqueue("https://example.com") // duplicate

	url1 := f.Dequeue()
	url2 := f.Dequeue() // should be empty

	assert.Equal(t, "https://example.com", url1)
	assert.Empty(t, url2)
}

func TestFrontier_EmptyDequeue(t *testing.T) {
	f := New()

	url := f.Dequeue()

	assert.Empty(t, url)
}
