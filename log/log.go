// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package log implements a simple logging package. It defines a type, Logger,
// with methods for formatting output. It also has a predefined 'standard'
// Logger accessible through helper functions Print[f|ln], Fatal[f|ln], and
// Panic[f|ln], which are easier to use than creating a Logger manually.
// That logger writes to standard error and prints the date and time
// of each logged message.
// Every log message is output on a separate line: if the message being
// printed does not end in a newline, the logger will add one.
// The Fatal functions call os.Exit(1) after writing the log message.
// The Panic functions call panic after writing the log message.
package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

type Level int

const (
	ERROR Level = iota + 1
	WARNING
	INFO
)

const LstdFlags = log.LstdFlags | log.Lmsgprefix | log.LUTC

var bold = color.New(color.Bold)

var tags = [4]string{
	bold.Add(color.FgRed).Sprint("[ERR]"),
	bold.Add(color.FgYellow).Sprint("[WRN]"),
	bold.Add(color.FgGreen).Sprint("[INF]"),
}

func prefix(level Level) string {
	return tags[level-1] + " "
}

func GetEnvLevel() Level {
	level := os.Getenv("LOG_LEVEL")
	switch level {
	case "INFO":
		return INFO
	case "ERROR":
		return ERROR
	default:
		return WARNING
	}
}

func GetLevelOrDefault(level Level) Level {
	switch level {
	case ERROR:
		return level
	case WARNING:
		return level
	default:
		return INFO
	}
}

func Writer(defaultWriter *os.File) *os.File {
	file := os.Getenv("LOG_FILE")
	value := strings.ToLower(file)
	if value == "false" || value == "" {
		return defaultWriter
	}

	writer, err := os.OpenFile(file, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.SetFlags(LstdFlags)
		log.SetPrefix(prefix(ERROR))
		log.Panic(err)
	}

	return writer
}

var std = log.New(Writer(os.Stdout), prefix(INFO), LstdFlags)

func Logger(level Level) *log.Logger {
	var writer io.Writer
	isDebug := strings.ToLower(os.Getenv("APP_DEBUG"))
	if level <= GetEnvLevel() || isDebug == "true" {
		if level == INFO {
			writer = Writer(os.Stdout)
		} else {
			writer = Writer(os.Stderr)
		}
	} else {
		writer = io.Discard
	}
	std := log.New(writer, prefix(level), log.LstdFlags|log.Lmsgprefix|log.LUTC)

	return std
}

func Output(calldepth int, level Level, format string, v ...any) {
	isDebug := strings.ToLower(os.Getenv("APP_DEBUG"))
	if level <= GetEnvLevel() || isDebug == "true" {
		format = fmt.Sprintf(format, v...)
		var writer io.Writer
		if level == INFO {
			writer = Writer(os.Stdout)
		} else {
			writer = Writer(os.Stderr)
		}
		std.SetOutput(writer)
		std.SetFlags(log.LstdFlags | log.Lmsgprefix | log.LUTC)
		std.SetPrefix(prefix(level))
		std.Output(calldepth, format)
	}
}

func Info(format string, v ...any) {
	Output(2, INFO, format, v...)
}

func Warning(format string, v ...any) {
	Output(2, WARNING, format, v...)
}

func Error(format string, v ...any) {
	Output(2, ERROR, format, v...)
}

func Debug(level Level, format string, v ...any) {
	Output(2, level, format, v...)
}
