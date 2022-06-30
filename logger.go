package main

/* Extremely Simple Logger Written By Osborn @chaos */
/* Testing this module will be difficult since the functions (mostly) just operate as impure functions */
/* I've debugged with a main function anyway */

/* This package can be used as a base for a more advanced logger but its quite sufficient */

/*
Api: Just two functions (methods)
LogC(msg)
LogF(msg, filename) err // Error is for potential file opening issues
*/

import (
	"fmt"
	"os"
	"time"
	"strings"
)

type LogLevel string

const (
	/* string values are good for debugging */
	_ = iota
	LogInfo = "Info"
	LogDebug = "Debug"
	LogWarning = "Warning"
	LogError = "Error"
)

type Logger struct {
	logLevel LogLevel
	strict bool	/* When not in strict mode loginfo messages to console will not appear */
}

func (l *Logger) Strict() {
	if l.strict == false {
		l.strict = true
	}
}

func (l *Logger) RemoveStrict() {
	if l.strict == true {
		l.strict = false
	}
}

func (l *Logger) SetLogLevel(level LogLevel) {
	l.logLevel = level
}

func (l Logger) log(level LogLevel, msg string) {
	msg = prettify(msg, l.logLevel)

	if l.strict {
		fmt.Fprintf(os.Stderr, msg)
		return
	} else {
		if level != LogInfo {
			fmt.Fprintf(os.Stderr, msg)
		}
	}
}

/* Output log message to console */
func (l Logger) LogC(msg string) {
	switch l.logLevel {
		case LogInfo:
			l.log(LogInfo, msg)
		case LogDebug:
			l.log(LogDebug, msg)
		case LogWarning:
			l.log(LogWarning, msg)
		case LogError:
			l.log(LogError, msg)
		default:
			printError("Invalid Log Level")
	}
}

func printError(msg string) {
	fmt.Printf("Logger_Error: ", msg)
}

/* Output log message to file */
func (l Logger) LogF(msg string, filename string) error {
	if filename == "" {
		filename = "file.log"
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	logMsg := prettify(msg, l.logLevel)

	file.Write([]byte(logMsg))

	return nil
}


func prettify(msg string, level LogLevel) string {
	val := string(level)
	return fmt.Sprintf("[%s - %s]: %s\n", strings.ToUpper(val), time.Now(), msg)
}

/*
func main() {
	l := Logger{
		logLevel: LogDebug,
	}

	l.Strict()

	err := l.LogF("Test Message", "")
	if err != nil {
		fmt.Println(err)
	}

	l.LogC("Test Message")
	l.SetLogLevel(LogInfo)
	l.RemoveStrict()

	l.LogF("Another Message", "")
}
*/
