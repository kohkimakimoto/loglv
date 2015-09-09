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

func IsDebug() bool {
	return writer.lv >= LvDebug
}

func IsInfo() bool {
	return writer.lv >= LvInfo
}

func IsWarning() bool {
	return writer.lv >= LvWarning
}

func IsError() bool {
	return writer.lv >= LvError
}

func IsQuiet() bool {
	return writer.lv >= LvQuiet
}

func Init() {
	log.SetOutput(writer)
}
