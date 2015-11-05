package loglv

import (
	"errors"
	"io"
	"log"
	"os"
)

const (
	// LvDebug is a debug level.
	LvDebug = 4
	// LvInfo is a info level.
	LvInfo = 3
	// LvWarning is a warning level.
	LvWarning = 2
	// LvError is a error level.
	LvError = 1
	// LvQuiet is a quiet level.
	LvQuiet = 0
)

var inited = false

type leveledWriter struct {
	lv  int
	out io.Writer
}

func (w leveledWriter) Write(p []byte) (n int, err error) {
	if w.lv != LvQuiet {
		n, err = writer.out.Write(p)
	}
	return
}

func (w *leveledWriter) LvToString() string {
	if w.lv == LvDebug {
		return "debug"
	}
	if w.lv == LvInfo {
		return "info"
	}
	if w.lv == LvWarning {
		return "warning"
	}
	if w.lv == LvError {
		return "error"
	}
	if w.lv == LvQuiet {
		return "quiet"
	}

	panic("Unknown level")
}

var writer = &leveledWriter{
	lv:  LvInfo,
	out: os.Stderr,
}

// SetLv sets log level.
func SetLv(level int) {
	writer.lv = level
}

// Lv returns log level
func Lv() int {
	return writer.lv
}

// Lv returns log level string
func LvString() string {
	return writer.LvToString()
}

// Writer returns a leveled log writer.
func Writer() io.Writer {
	return writer
}

// SetLevelByString sets a log level by string.
func SetLevelByString(lv string) error {
	if lv == "debug" {
		writer.lv = LvDebug
	} else if lv == "info" {
		writer.lv = LvInfo
	} else if lv == "warning" {
		writer.lv = LvWarning
	} else if lv == "error" {
		writer.lv = LvError
	} else if lv == "quiet" {
		writer.lv = LvQuiet
	} else {
		return errors.New("Invalid log level '" + lv + "'.")
	}
	return nil
}

// SetOutput sets writer that is used by logging.
func SetOutput(w io.Writer) {
	writer.out = w
}

// IsDebug returns true value if the level is debug.
func IsDebug() bool {
	return writer.lv >= LvDebug
}

// IsInfo returns true value if the level is info.
func IsInfo() bool {
	return writer.lv >= LvInfo
}

// IsWarning returns true value if the level is warning.
func IsWarning() bool {
	return writer.lv >= LvWarning
}

// IsError returns true value if the level is error.
func IsError() bool {
	return writer.lv >= LvError
}

// IsQuiet returns true value if the level is quiet.
func IsQuiet() bool {
	return writer.lv >= LvQuiet
}

// Init initializes logger output.
func Init() {
	if inited {
		return
	}
	log.SetOutput(writer)
	inited = true
}
