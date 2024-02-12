package env

import (
	"bytes"
	"flag"
	"io"
	"os"
	"strings"

	"github.com/chocokacang/chocokacang/log"
	"github.com/chocokacang/chocokacang/utils"
)

type envMap map[string]string

func Init() {
	envMap := read()
	for key, value := range envMap {
		os.Setenv(key, value)
	}
}

func read() envMap {
	filename := flag.String("env", ".env", "Path to environment file")
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		log.Debug(log.WARNING, "Failed to %v", err)
		return nil
	}
	defer file.Close()

	return Parse(file)
}

func Parse(r io.Reader) envMap {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		log.Debug(log.WARNING, "Failed to parse env variable. Got error: %v", err)
	}

	parsed, _ := utils.UnmarshalBytes(buf.Bytes())

	return parsed
}

// UnmarshalBytes parses env file from byte slice of chars, returning a map of keys and values.
func UnmarshalBytes(src []byte) (map[string]string, error) {
	out := make(map[string]string)
	err := parseBytes(src, out)

	return out, err
}

func GetString(key string, s string) string {
	val := os.Getenv(key)
	if val == "" {
		return s
	}
	return val
}

func GetBool(key string, b bool) bool {
	val := os.Getenv(key)
	varLower := strings.ToLower(val)
	if val == "" {
		return b
	}
	return varLower == "true"
}
