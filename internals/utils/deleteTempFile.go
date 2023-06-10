package utils

import (
	"os"
)

func DeleteTempFile(path string) error {
	err := os.Remove(path)
	return err
}
