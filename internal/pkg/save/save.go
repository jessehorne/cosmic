package save

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// GetSavePath gets the path to the OS-independent user config directory, or an error
// if it doesn't exist.
func GetSavePath() (string, error) {
	baseDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return baseDir + "/cosmic", nil
}

// CreateSaveIfNotExists creates the save / project directory in the OS-independent
// user config save path for cosmic if it doesn't already exist.
func CreateSaveIfNotExists() error {
	savePath, err := GetSavePath()
	if err != nil {
		return err
	}

	err = copyDir("data/save/", savePath)
	if err != nil {
		return err
	}

	return nil
}

// copyDir copies the contents of a directory recursively to the destination
func copyDir(src, dst string) error {
	if err := os.MkdirAll(dst, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory %s: %w", dst, err)
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("failed to read source directory %s: %w", src, err)
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		fmt.Println(srcPath, dstPath)

		if entry.IsDir() {
			if cpErr := copyDir(srcPath, dstPath); cpErr != nil {
				return cpErr
			}
		} else {
			// Copy the file
			if err := copyFile(srcPath, dstPath); err != nil {
				log.Printf("Warning: Could not copy file %s: %v", srcPath, err)
				continue // Log the warning and continue with the next file
			}
		}
	}
	return nil
}

// copyFile copies the contents of the file at src to the file at dst.
func copyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}
