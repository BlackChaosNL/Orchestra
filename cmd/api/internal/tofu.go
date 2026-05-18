package internal

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/opentofu/tofu-exec/tfexec"
	"github.com/opentofu/tofudl"
)

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
		fmt.Errorf("Error when writing the file %s: %s", execPath, err)
	}

	return execPath
}

func GetTofu(tofuVersion string) (*tfexec.Tofu, string, string) {
	_, tofuDir := PrepareTemporaryDictionary("go-orchestra-opentofu-")
	_, workPath := PrepareTemporaryDictionary("go-orchestra-opentofu-work-dir-")
	execPath := DownloadTofu(tofuVersion, tofuDir)
	tofu, err := tfexec.NewTofu(workPath, execPath)

	if err != nil {
		fmt.Errorf("Error running Tofu: %s", err)
	}

	return tofu, tofuDir, workPath
}
