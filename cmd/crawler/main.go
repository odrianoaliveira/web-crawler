package main

import (
	"flag"
	"fmt"
	"github.com/odrianoaliveira/web-crawler/internal/frontier"
	"github.com/odrianoaliveira/web-crawler/internal/seed"
)

func main() {
	fmt.Println("Starting the web crawler!")
	seedFile := flag.String("seed-url", "urls.txt", "Path to the seed URL file")
	flag.Parse()

	err := run(seedFile)
	if err != nil {
		fmt.Printf("Crawler rrror: %v\n", err)
	}
}

func run(seedFile *string) error {
	fmt.Println("Seed URL file:", *seedFile)

	urls, err := seed.ReadURLs(*seedFile)
	if err != nil {
		return fmt.Errorf("Error parsing URLs: %v\n", err)
	}

	f := frontier.New()
	for _, url := range urls {
		fmt.Println("Seed URL:", url)
		f.Enqueue(url)
	}

	for !f.IsEmpty() {
		url := f.Dequeue()
		fmt.Println("Visit:", url)
	}

	return nil
}
