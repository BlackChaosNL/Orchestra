package internal

import (
	"fmt"
	"os"
)

func PrepareTemporaryDictionary(dir string) (error, string) {
	tempDir, err := os.MkdirTemp("", dir)
	if err != nil {
		return fmt.Errorf("Can not create a temporary folder... %s", err), ""
	}
	return nil, tempDir
}

func RemoveFolder(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return fmt.Errorf("Can not remove the temporary folder... %s", err)
	}
	return nil
}
