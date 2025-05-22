package seed

import (
	"bufio"
	"os"
	"strings"
)

func ReadURLs(filepath string) (urls []string, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || line[0] == '#' {
			continue // Skip empty lines and comments
		}
		urls = append(urls, line)

		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}

	return urls, nil
}
