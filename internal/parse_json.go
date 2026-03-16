package internal

import (
	"os"

	"github.com/kataras/golog"
	"github.com/tidwall/gjson"
)

var file_content []byte

func ReadEnvFile(p string) {
	file_content, err := os.ReadFile(p)

	if err != nil {
		golog.Info("Env file read requested, error occurred. Skipping...", file_content)
	}
}

func GetFromEnvFile(query string) string {
	// INFO: To query, use: https://github.com/tidwall/gjson syntax.
	return gjson.Get(string(file_content), query).String()
}
