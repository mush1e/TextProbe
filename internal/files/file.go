package files

import (
	"os"
	"path/filepath"
)

func TraverseDirectory(dirPath string, processFile func(filePath string) error) error {
	return filepath.WalkDir(dirPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		return processFile(path)
	})
}

func ReadFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
