package config

import (
	"os"

	"github.com/chocokacang/chocokacang/env"
	"github.com/chocokacang/chocokacang/log"
)

type Config struct {
	AppName  string
	AppEnv   string
	AppDebug bool
	AppPort  string
	DbHost   string
	DbPort   string
	DbUser   string
	DbPass   string
	LogFile  string
	LogLevel log.Level
}

var Get = &Config{
	AppName:  os.Getenv("APP_NAME"),
	AppDebug: env.GetBool("APP_DEBUG", false),
	AppPort:  env.GetString("APP_PORT", "8080"),
	DbHost:   os.Getenv("DB_HOST"),
	DbPort:   os.Getenv("DB_PORT"),
	DbUser:   os.Getenv("DB_USER"),
	DbPass:   os.Getenv("DB_PASS"),
	LogFile:  os.Getenv("LOG_FILE"),
	LogLevel: log.GetEnvLevel(),
}
