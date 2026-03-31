package config

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/kataras/golog"
	"github.com/opentofu/tofu-exec/tfexec"
	"github.com/opentofu/tofudl"
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
	if err == nil {
		golog.Fatalf("Can not remove the temporary folder... %s", err)
	}
}

func DownloadTofu(tofuVersion string, dir string) string {
	opts := tofudl.DownloadOptVersion(tofudl.Version(tofuVersion))
	dl, err := tofudl.New()
	if err != nil {
		log.Fatalf("Error when instantiating tofudl %s", err)
	}
	binary, _ := dl.Download(context.TODO(), opts)
	execPath := filepath.Join(dir, "tofu")
	// Windows executable case
	if runtime.GOOS == "windows" {
		execPath += ".exe"
	}
	if err := os.WriteFile(execPath, binary, 0755); err != nil {
		golog.Fatalf("Error when writing the file %s: %s", execPath, err)
	}

	return execPath
}

func GetTofu(tofuVersion string) (*tfexec.Tofu, string, string) {
	tofuDir := PrepareTemporaryDictionary("go-orchestra-opentofu-")
	workPath := PrepareTemporaryDictionary("go-orchestra-opentofu-work-dir-")
	execPath := DownloadTofu(tofuVersion, tofuDir)
	tofu, err := tfexec.NewTofu(workPath, execPath)

	if err != nil {
		golog.Fatalf("Error running Tofu: %s", err)
	}

	return tofu, tofuDir, workPath
}
