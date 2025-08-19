package utils

import (
	"os"
	"strings"
)

func RemoveIfExists(filePath string) error {
	if _, err := os.Stat(filePath); err == nil {
		return os.Remove(filePath)
	} else if !os.IsNotExist(err) {
		return err
	}
	return nil
}

// suffix: .dot, .elf, .bin
func AddSuffix(raw, suffix string) string {
	if strings.HasSuffix(raw, suffix) {
		return raw
	}
	return raw + suffix
}
