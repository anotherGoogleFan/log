package log

import (
	"fmt"
	"os"

	rus "github.com/sirupsen/logrus"
)

// Fields is used to define structured records for a log entry.
type Fields rus.Fields

func (d Fields) print(f logFunc, skip int, v []interface{}) {
	var msg interface{} = v
	if len(v) == 1 {
		msg = v[0]
	}
	d["pos"] = getFilePos(skip + 1)
	if GetMode() == Production {
		d["process"] = os.Args[0]
		d["release"] = std.GetRelease()
	}
	std.print(d, msg, f)
}

func (d Fields) printf(f logFunc, skip int, format string, v []interface{}) {
	d.print(f, skip+1, []interface{}{fmt.Sprintf(format, v...)})
}

type logFunc func(entry *rus.Entry, args ...interface{})

var (
	fPanic = (*rus.Entry).Panic
	fFatal = (*rus.Entry).Fatal
	fError = (*rus.Entry).Error
	fWarn  = (*rus.Entry).Warn
	fInfo  = (*rus.Entry).Info
	fDebug = (*rus.Entry).Debug
)

// Panic logs at the panic level and then panic.
func (d Fields) Panic(v ...interface{}) {
	d.print(fPanic, 1, v)
}

// Fatal logs at the fatal level and then os.Exit.
func (d Fields) Fatal(v ...interface{}) {
	d.print(fFatal, 1, v)
}

// Error logs at the error level.
func (d Fields) Error(v ...interface{}) {
	d.print(fError, 1, v)
}

// Warn logs at the warn level.
func (d Fields) Warn(v ...interface{}) {
	d.print(fWarn, 1, v)
}

// Info logs at the info level.
func (d Fields) Info(v ...interface{}) {
	d.print(fInfo, 1, v)
}

// Debug logs at the debug level.
func (d Fields) Debug(v ...interface{}) {
	d.print(fDebug, 1, v)
}

// Panicf is the "format" version of Panic.
func (d Fields) Panicf(format string, v ...interface{}) {
	d.printf(fPanic, 1, format, v)
}

// Fatalf is the "format" version of Fatal.
func (d Fields) Fatalf(format string, v ...interface{}) {
	d.printf(fFatal, 1, format, v)
}

// Errorf is the "format" version of Error.
func (d Fields) Errorf(format string, v ...interface{}) {
	d.printf(fError, 1, format, v)
}

// Warnf is the "format" version of Warn.
func (d Fields) Warnf(format string, v ...interface{}) {
	d.printf(fWarn, 1, format, v)
}

// Infof is the "format" version of Info.
func (d Fields) Infof(format string, v ...interface{}) {
	d.printf(fInfo, 1, format, v)
}

// Debugf is the "format" version of Debug.
func (d Fields) Debugf(format string, v ...interface{}) {
	d.printf(fDebug, 1, format, v)
}

// Panic logs at the panic level and then panic.
func (d Fields) PanicWithSkip(skip int, v ...interface{}) {
	d.print(fPanic, skip+1, v)
}

// Fatal logs at the fatal level and then os.Exit.
func (d Fields) FatalWithSkip(skip int, v ...interface{}) {
	d.print(fFatal, skip+1, v)
}

// Error logs at the error level.
func (d Fields) ErrorWithSkip(skip int, v ...interface{}) {
	d.print(fError, skip+1, v)
}

// Warn logs at the warn level.
func (d Fields) WarnWithSkip(skip int, v ...interface{}) {
	d.print(fWarn, skip+1, v)
}

// Info logs at the info level.
func (d Fields) InfoWithSkip(skip int, v ...interface{}) {
	d.print(fInfo, skip+1, v)
}

// Debug logs at the debug level.
func (d Fields) DebugWithSkip(skip int, v ...interface{}) {
	d.print(fDebug, skip+1, v)
}
