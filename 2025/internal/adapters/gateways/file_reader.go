package gateways

import (
	"bufio"
	"fmt"
	"os"
)

// SimpleFileReader reads files from the filesystem
type SimpleFileReader struct{}

// NewSimpleFileReader creates a new file reader
func NewSimpleFileReader() *SimpleFileReader {
	return &SimpleFileReader{}
}

// ReadLines reads a file and returns its contents as a slice of lines
func (r *SimpleFileReader) ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fmt.Println("finished scanning")
	return lines, scanner.Err()
}
