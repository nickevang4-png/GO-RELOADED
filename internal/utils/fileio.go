package utils

import "os"

// ReadFile reads a text file and returns its contents.
func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// WriteFile writes a string into a file (overwrites if exists).
func WriteFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}
