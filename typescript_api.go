package web

import (
	"fmt"
	"os"
	"path/filepath"
)

func openOrCreateFile(path string) (*os.File, error) {
	// Create all directories along the path if they don't exist
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create directories: %v", err)
	}

	// Open the file with O_CREATE and O_TRUNC flags
	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open or create file: %v", err)
	}

	return file, nil
}
