package log

import (
	"fmt"
	"io"
	"reflect"

	rus "github.com/sirupsen/logrus"
)

// Level for logging
type Level int

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel = Level(rus.PanicLevel)
	// FatalLevel level. Logs and then calls `os.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel = Level(rus.FatalLevel)
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel = Level(rus.ErrorLevel)
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel = Level(rus.WarnLevel)
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel = Level(rus.InfoLevel)
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel = Level(rus.DebugLevel)
)

func (l Level) String() string {
	switch l {
	case PanicLevel:
		return "panic"
	case FatalLevel:
		return "fatal"
	case ErrorLevel:
		return "error"
	case WarnLevel:
		return "warn"
	case InfoLevel:
		return "info"
	case DebugLevel:
		return "debug"
	}
	return ""
}

// SetLevel sets the log level. The argument level can be either string or Level
// type.
func SetLevel(level interface{}) error {
	switch l := level.(type) {
	case Level:
		std.SetLevel(l)
	case string:
		lev, err := parseLevel(l)
		if err != nil {
			return err
		}
		std.SetLevel(lev)
	default:
		return fmt.Errorf("unknown log type: %v, only support int or string type", reflect.TypeOf(level))
	}
	return nil
}
func parseLevel(s string) (Level, error) {
	l, err := rus.ParseLevel(s)
	return Level(l), err
}

// GetLevel returns the current log level.
func GetLevel() Level {
	return std.GetLevel()
}

// Mode of execution of the program: eithor Develop of Production mode.
type Mode int

const (
	// Develop mode for daily development.
	Develop Mode = iota
	// Production mode for online production.
	Production
)

func (m Mode) String() string {
	switch m {
	case Develop:
		return "develop"
	case Production:
		return "production"
	}
	return ""
}

// SetMode set "develop" or "production" mode for the logging.
// argument mode could be:
// 1. log.Production or "production"
// 2. log.Develop or "develop"
func SetMode(mode interface{}) error {
	switch m := mode.(type) {
	case Mode:
		std.SetMode(m)
	case string:
		mode, err := parseMode(m)
		if err != nil {
			return err
		}
		std.SetMode(mode)
	default:
		return fmt.Errorf("unknown log mode: %v, only support int or string type", reflect.TypeOf(mode))
	}
	return nil
}
func parseMode(mode string) (Mode, error) {
	switch mode {
	case "develop":
		return Develop, nil
	case "production":
		return Production, nil
	}
	return 0, fmt.Errorf("not a valid Mode: %q", mode)
}

// GetMode returns the current Mode.
func GetMode() Mode {
	return std.GetMode()
}

// SetRelease sets the release version of the current program for tracing back
// from log to source code.
func SetRelease(release string) {
	std.SetRelease(release)
}

// GetRelease returns the current release version.
func GetRelease() string {
	return std.GetRelease()
}

// SetFormatter sets the standard logger formatter.
// Options: logstash, json, text.
func SetFormatter(formatter string) {
	std.SetFormatter(formatter)
}

func GetFormatter() string {
	return std.GetFormatter()
}

// SetOutput sets the output of the std logger.
func SetOutput(o io.Writer) {
	std.SetOutput(o)
}

func SetLogrusLogger(l *rus.Logger) {
	std.SetLogrusLogger(l)
}

func SetLogrusFormatter(formatter rus.Formatter) {
	std.SetLogrusFormatter(formatter)
}
