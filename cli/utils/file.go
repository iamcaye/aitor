package utils

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/google/brotli/go/cbrotli"
)

// read file reads the contents of a file and returns it as a string.
func ReadFile(filePath string) (string, error) {
	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", fmt.Errorf("file does not exist: %s", filePath)
	}

	// Read the file contents
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return string(data), nil
}

// return files
func ReadDir(dirPath string) ([]os.FileInfo, error) {
	// Check if the directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("directory does not exist: %s", dirPath)
	}

	// Read the directory contents
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}
	var filesInfo []os.FileInfo
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			return nil, fmt.Errorf("error getting file info: %w", err)
		}
		filesInfo = append(filesInfo, info)
	}

	return filesInfo, nil
}

func CompressContent(content []byte) ([]byte, error) {
	var compressed bytes.Buffer
	start := time.Now()

	var writer cbrotli.Writer = *cbrotli.NewWriter(&compressed, cbrotli.WriterOptions{
		Quality: 11,
	})

	defer writer.Close()
	_, err := writer.Write([]byte(content))
	if err != nil {
		return nil, fmt.Errorf("error writing file: %w", err)
	}

	err = writer.Flush()
	if err != nil {
		return nil, fmt.Errorf("error flushing file: %w", err)
	}

	// Close the writer to ensure all data is written
	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("error closing writer: %w", err)
	}

	// Log the time taken for compression
	duration := time.Since(start)
	fmt.Printf("Compression took: %s\n", duration)
	return compressed.Bytes(), nil
}

func DecompressFile(compressedContent []byte) (string, error) {
	var decompressed bytes.Buffer
	start := time.Now()

	reader := cbrotli.NewReader(bytes.NewReader(compressedContent))
	defer reader.Close()

	_, err := decompressed.ReadFrom(reader)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	// Log the time taken for decompression
	duration := time.Since(start)
	fmt.Printf("Decompression took: %s\n", duration)
	return decompressed.String(), nil
}
