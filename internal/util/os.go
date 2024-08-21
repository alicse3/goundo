package util

import (
	"io"
	"os"
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
