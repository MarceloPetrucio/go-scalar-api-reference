package scalar

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// ensureFileURL takes a file path and returns a complete URL with the file:// scheme
func ensureFileURL(filePath string) (string, error) {
	// Checks if the path already has the file:// prefix
	if strings.HasPrefix(filePath, "file://") {
		// Checks if the remainder of the path is absolute
		if path := strings.TrimPrefix(filePath, "file://"); !filepath.IsAbs(path) {
			// If not absolute, resolves relative to the current working directory
			currentDir, err := os.Getwd()
			if err != nil {
				return "", fmt.Errorf("error getting current directory: %w", err)
			}
			resolvedPath := filepath.Join(currentDir, path)
			return "file://" + resolvedPath, nil
		}
		// If it is already absolute, returns as is
		return filePath, nil
	}

	// If the path does not start with file://, checks if it is absolute
	if filepath.IsAbs(filePath) {
		// If it is absolute, just adds the file:// prefix
		return "file://" + filePath, nil
	}

	// If it is a relative path, resolves relative to the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current directory: %w", err)
	}
	resolvedPath := filepath.Join(currentDir, filePath)
	return "file://" + resolvedPath, nil
}

// fetchContentFromURL reads the content from a URL and returns it as a string
func fetchContentFromURL(fileURL string) (string, error) {
	resp, err := http.Get(fileURL)
	if err != nil {
		return "", fmt.Errorf("error getting file content: %w", err)
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading file content: %w", err)
	}

	return string(content), nil
}

func readFileFromURL(fileURL string) ([]byte, error) {
	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}

	if parsedURL.Scheme != "file" {
		return nil, fmt.Errorf("unsupported URL scheme: %s", parsedURL.Scheme)
	}

	// Reads the file from the path specified in the URL
	return os.ReadFile(parsedURL.Path)
}
