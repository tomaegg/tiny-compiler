package utils

import "os"

func RemoveIfExists(filePath string) error {
	if _, err := os.Stat(filePath); err == nil {
		return os.Remove(filePath)
	} else if !os.IsNotExist(err) {
		return err
	}
	return nil
}
