package internal

import (
	"os"

	"github.com/kataras/golog"
)

func PrepareTemporaryDictionary(dir string) string {
	tempDir, err := os.MkdirTemp("", dir)
	if err != nil {
		golog.Fatalf("Can not create a temporary folder... %s", err)
	}
	return tempDir
}

func RemoveFolder(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		golog.Fatalf("Can not remove the temporary folder... %s", err)
	}
}
