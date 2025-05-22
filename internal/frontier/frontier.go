package frontier

import (
	"container/list"
	"sync"
)

type Frontier struct {
	mu      sync.Mutex
	queue   *list.List
	visited map[string]bool
}

func New() *Frontier {
	return &Frontier{
		queue:   list.New(),
		visited: make(map[string]bool),
	}
}

func (f *Frontier) Enqueue(url string) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if !f.visited[url] {
		f.queue.PushBack(url)
		f.visited[url] = true
	}
}

func (f *Frontier) Dequeue() string {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.queue.Len() == 0 {
		return ""
	}

	url := f.queue.Front().Value.(string)
	f.queue.Remove(f.queue.Front())
	return url
}

func (f *Frontier) IsEmpty() bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	return f.queue.Len() == 0
}
