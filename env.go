package chocokacang

import (
	"bytes"
	"io"
	"os"

	"github.com/chocokacang/chocokacang/log"
	"github.com/chocokacang/chocokacang/utils"
)

type envMap map[string]string

func init() {
	envMap := readEnvFile()
	for key, value := range envMap {
		os.Setenv(key, value)
	}
}

func readEnvFile() envMap {
	file, err := os.Open(".env")
	if err != nil {
		log.Debug(log.WARNING, "Failed to %v", err)
		return nil
	}
	defer file.Close()

	return parseEnvFile(file)
}

func parseEnvFile(r io.Reader) envMap {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		log.Debug(log.WARNING, "Failed to parse env variable. Got error: %v", err)
	}

	parsed, _ := utils.UnmarshalBytes(buf.Bytes())

	return parsed
}
