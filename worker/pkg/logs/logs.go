package logs

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

const (
	MinLogLevel = LevelDebug

	LevelDebug   uint8 = 0
	LevelInfo    uint8 = 1
	LevelWarning uint8 = 2
	LevelError   uint8 = 3
	LevelFatal   uint8 = 4

	MaxLogLevel = LevelFatal
)

var minLogLevel = LevelInfo

func SetMinLevel(level uint8) error {
	if level < LevelDebug || level > LevelFatal {
		return fmt.Errorf(
			"invalid log level: value=%d, min=%d, max=%d",
			level, LevelDebug, LevelFatal,
		)
	}

	minLogLevel = level
	return nil
}

func SetFileOutput(fname string) error {
	f, err := os.OpenFile(
		fname, syscall.O_CREAT|syscall.O_APPEND|syscall.O_WRONLY, 0o600,
	)
	if err != nil {
		return fmt.Errorf("os.OpenFile() failed: cannot init logger: %w", err)
	}

	log.SetOutput(f)
	return nil
}

const (
	levelDebugStr   = "DEBUG"
	levelInfoStr    = "INFO"
	levelWarningStr = "WARNING"
	levelErrorStr   = "ERROR"
)

func Debug(v ...interface{}) {
	printLog(LevelDebug, levelDebugStr, v...)
}

func DebugF(format string, v ...interface{}) {
	printLogF(LevelDebug, levelDebugStr, format, v...)
}

func Info(v ...interface{}) {
	printLog(LevelInfo, levelInfoStr, v...)
}

func InfoF(format string, v ...interface{}) {
	printLogF(LevelInfo, levelInfoStr, format, v...)
}

func Warning(v ...interface{}) {
	printLog(LevelWarning, levelWarningStr, v...)
}

func WarningF(format string, v ...interface{}) {
	printLogF(LevelWarning, levelWarningStr, format, v...)
}

func Error(v ...interface{}) {
	printLog(LevelError, levelErrorStr, v...)
}

func ErrorF(format string, v ...interface{}) {
	printLogF(LevelError, levelErrorStr, format, v...)
}

func printLog(
	lvl uint8, slvl string, v ...interface{},
) {
	if lvl < minLogLevel {
		return
	}

	xv := makeLogFnArgs(slvl, v)
	log.Println(xv...)
}

func printLogF(
	lvl uint8, slvl string, format string, v ...interface{},
) {
	if lvl < minLogLevel {
		return
	}

	xv := makeLogFnArgs(slvl, v)
	log.Printf("%s "+format, xv...)
}

func makeLogFnArgs(slvl string, v []interface{}) []interface{} {
	xv := make([]interface{}, 0, len(v)+1)

	xv = append(xv, fmt.Sprintf("[%s]", slvl))
	xv = append(xv, v...)

	return xv
}
