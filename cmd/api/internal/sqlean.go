package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/BlackChaosNL/Orchestra/config"
	"golift.io/xtractr"
)

func buildDownloadString() (string, string) {
	version := "0.28.2"
	file := fmt.Sprintf("sqlean-%s-x64", config.GetStrFromEnv("GOOS", "linux"))
	return fmt.Sprintf("https://github.com/nalgeon/sqlean/releases/tag/%s/%s.zip", version, file), file
}

func downloadZip(url string, dir string, file_name string) error {
	_, temp_dir := PrepareTemporaryDictionary(dir)

	file, err := os.Create(fmt.Sprintf("%s/%s.zip", temp_dir, file_name))
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Check if the server returned a successful status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// 5. Stream the response body directly to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save file content: %w", err)
	}

	return nil
}

func GetLatestSqleanRelease() (bool, string, string) {
	dir := "go-orchestra-sqlean-"
	download_url, file_name := buildDownloadString()

	downloadZip(download_url, dir, file_name)

	if _, files, _, err := xtractr.ExtractFile(&xtractr.XFile{
		FilePath:  fmt.Sprintf("%s/%s.zip", dir, file_name),
		OutputDir: fmt.Sprintf("%s/%s", dir, file_name),
	}); err != nil || files == nil {
		return false, err.Error(), ""
	}

	return true, dir, file_name
}
