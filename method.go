package log

// Panic without custom Fields.
func Panic(v ...interface{}) {
	Fields{}.print(fPanic, 1, v)
}

// Fatal without custom Fields.
func Fatal(v ...interface{}) {
	Fields{}.print(fFatal, 1, v)
}

// Error without custom Fields.
func Error(v ...interface{}) {
	Fields{}.print(fError, 1, v)
}

// Warn without custom Fields.
func Warn(v ...interface{}) {
	Fields{}.print(fWarn, 1, v)
}

// Info without custom Fields.
func Info(v ...interface{}) {
	Fields{}.print(fInfo, 1, v)
}

// Debug without custom Fields.
func Debug(v ...interface{}) {
	Fields{}.print(fDebug, 1, v)
}

// Panicf without custom Fields.
func Panicf(format string, v ...interface{}) {
	Fields{}.printf(fPanic, 1, format, v)
}

// Fatalf without custom Fields.
func Fatalf(format string, v ...interface{}) {
	Fields{}.printf(fFatal, 1, format, v)
}

// Errorf without custom Fields.
func Errorf(format string, v ...interface{}) {
	Fields{}.printf(fError, 1, format, v)
}

// Warnf without custom Fields.
func Warnf(format string, v ...interface{}) {
	Fields{}.printf(fWarn, 1, format, v)
}

// Infof without custom Fields.
func Infof(format string, v ...interface{}) {
	Fields{}.printf(fInfo, 1, format, v)
}

// Debugf without custom Fields.
func Debugf(format string, v ...interface{}) {
	Fields{}.printf(fDebug, 1, format, v)
}

// CheckFatal calls Fatal(err) if err != nil.
func CheckFatal(err error) {
	if err != nil {
		Fields{}.print(fFatal, 1, []interface{}{err})
	}
}
