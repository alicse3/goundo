package util

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// MoveFile moves the file from source to destination location
func MoveFile(src string, dst string) error {
	// Open the source file for reading
	srcFile, err := os.OpenFile(src, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file for writing
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Copy the contents of the source file to the destination file
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	// Remove the source file
	if err := os.Remove(src); err != nil {
		return err
	}

	return nil
}

// MoveDirectory moves the directory and it contents from source to destination location
func MoveDirectory(src string, dst string) error {
	if err := filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		// Construct relative path
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return nil
		}
		dstPath := filepath.Join(dst, relPath)

		if d.IsDir() {
			// Create the directory in the destination path
			if err := os.MkdirAll(dstPath, os.ModePerm); err != nil {
				return err
			}
		} else {
			// Move the file to the destination path
			if err := MoveFile(path, dstPath); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	// Remove the source file
	if err := os.RemoveAll(src); err != nil {
		return err
	}

	return nil
}
