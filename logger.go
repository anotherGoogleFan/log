package log

import (
	"io"
	"sync"

	"github.com/sirupsen/logrus"
)

const (
	FORMAT_JSON = "json"
	FORMAT_TEXT = "text"
)

var std = newLogger()

type logger struct {
	formatter string
	mode      Mode
	release   string
	l         *logrus.Logger
	mu        sync.Mutex
}

func newLogger() *logger {
	return &logger{l: logrus.New()}
}

func NewLogger() *logger {
	return newLogger()
}

func (l *logger) print(fields Fields, msg interface{}, f logFunc) {
	l.mu.Lock()
	defer l.mu.Unlock()
	f(l.l.WithFields(logrus.Fields(fields)), msg)
}

func (l *logger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.l.Level = logrus.Level(level)
}

func (l *logger) GetLevel() Level {
	l.mu.Lock()
	defer l.mu.Unlock()
	return Level(l.l.Level)
}

func (l *logger) SetMode(mode Mode) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.mode = mode
}

func (l *logger) GetMode() Mode {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.mode
}

func (l *logger) SetRelease(release string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.release = release
}

func (l *logger) GetRelease() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.release
}

func (l *logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.l.Out = w
}

func (l *logger) SetLogrusLogger(rl *logrus.Logger) {
	l.mu.Lock()
	l.l = rl
	l.mu.Unlock()
}

func (l *logger) SetFormatter(formatter string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	switch formatter {
	case FORMAT_JSON:
		l.l.Formatter = new(logrus.JSONFormatter)
		l.formatter = formatter
	case FORMAT_TEXT:
		l.l.Formatter = new(logrus.TextFormatter)
		l.formatter = formatter
	default:
		l.l.Formatter = new(logrus.TextFormatter)
		l.formatter = FORMAT_TEXT
	}
}

func (l *logger) SetLogrusFormatter(formatter logrus.Formatter) {
	l.mu.Lock()
	l.l.Formatter = formatter
	l.mu.Unlock()
}

func (l *logger) GetFormatter() string {
	return l.formatter
}
