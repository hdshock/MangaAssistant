package utils

import (
	"os"
	"path/filepath"
)

// Check if a file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// Get the list of CBZ files in a directory
func GetCBZFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".cbz" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
} 