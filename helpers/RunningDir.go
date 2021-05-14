package helpers

import (
	"os"
	"path/filepath"
)

func RunningDir() (string, error) {
	file, err := os.Executable()

	if err != nil {
		return "", err
	}

	return filepath.Dir(file), nil
}
