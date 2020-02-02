/**
 * @Description: 
 * @Version: 1.0.0
 * @Author: liteng
 * @Date: 2020-02-02 14:32
 */

package common

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

const (
	Red    = "1;31"
	Green  = "1;32"
	Yellow = "1;33"
	Blue   = "1;34"
	Pink   = "1;35"
	Cyan   = "1;36"
)

func Color(code, msg string) string {
	return fmt.Sprintf("\033[%sm%s\033[m", code, msg)
}

const (
	debugLog uint8 = iota
	infoLog
	warnLog
	errorLog
	fatalLog
	disableLog
)

var (
	levels = []string{
		debugLog:   Color(Blue, "[DBG]"),
		infoLog:    Color(Pink, "[INF]"),
		warnLog:    Color(Yellow, "[WRN]"),
		errorLog:   Color(Red, "[ERR]"),
		fatalLog:   Color(Cyan, "[FAT]"),
		disableLog: "DISABLED",
	}
	Stdout = os.Stdout

	storeOnce sync.Once
	logger    *Logger
)

const (
	calldepth = 1
)

func levelName(level uint8) string {
	if int(level) >= len(levels) {
		return fmt.Sprintf("LEVEL%d", level)
	}
	return levels[int(level)]
}

func goid() string {
	var buf [18]byte
	n := runtime.Stack(buf[:], false)
	fields := strings.Fields(string(buf[:n]))
	if len(fields) <= 1 {
		return ""
	}
	return fields[1]
}

type Logger struct {
	level  uint8
	writer io.Writer
	logger *log.Logger
}

func NewLogger(outputPath string, level uint8, maxPerLogSizeMb, maxLogFolderSizeMb int64) *Logger {
	var perLogFileSize = defaultMaxFileSize
	var logFolderSize = defaultMaxFolderSize

	if maxPerLogSizeMb != 0 {
		perLogFileSize = maxPerLogSizeMb * MBSize
	}
	if maxLogFolderSizeMb != 0 {
		logFolderSize = maxLogFolderSizeMb * MBSize
	}

	fileWriter := NewFileWriter(outputPath, perLogFileSize, logFolderSize)
	logWriter := io.MultiWriter(os.Stdout, fileWriter)

	return &Logger{
		level:  level,
		writer: logWriter,
		logger: log.New(logWriter, "", log.Ldate|log.Lmicroseconds),
	}
}

func NewDefault(path string) *Logger {
	return NewLogger(path, 0, 20, 3)
}

func SingleStore() *Logger {
	storeOnce.Do(func() {
		path := GetConfig().GetValue("log", "path")
		logger = NewDefault(path)
	})
	return logger
}

func (l *Logger) Writer() io.Writer {
	return l.writer
}

func (l *Logger) Output(level uint8, a ...interface{}) {
	if l.level <= level {
		a = append([]interface{}{levelName(level), "GID", goid() + ","}, a...)
		l.logger.Output(calldepth, fmt.Sprintln(a...))
	}
}

func (l *Logger) Outputf(level uint8, format string, a ...interface{}) {
	if l.level <= level {
		a = append([]interface{}{levelName(level), "GID", goid() + ","}, a...)
		l.logger.Output(calldepth, fmt.Sprintf("%s %s %s "+format+"\n", a...))
	}
}

func (l *Logger) Debug(a ...interface{}) {
	if l.level > debugLog {
		return
	}

	pc, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		return
	}

	fn := runtime.FuncForPC(pc)
	a = append([]interface{}{fn.Name(), filepath.Base(file) + "@" + strconv.Itoa(line) + ":"}, a...)

	l.Output(debugLog, a...)
}

func (l *Logger) Debugf(format string, a ...interface{}) {
	if l.level > debugLog {
		return
	}

	pc, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		return
	}

	fn := runtime.FuncForPC(pc)
	a = append([]interface{}{fn.Name(), filepath.Base(file) + "@" + strconv.Itoa(line) + ":"}, a...)

	l.Outputf(debugLog, "%s %s "+format, a...)
}

func (l *Logger) Info(a ...interface{}) {
	pc, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		return
	}

	fn := runtime.FuncForPC(pc)
	a = append([]interface{}{fn.Name(), filepath.Base(file) + "@" + strconv.Itoa(line) + ":"}, a...)

	l.Output(infoLog, a...)
}

func (l *Logger) Infof(format string, a ...interface{}) {
	pc, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		return
	}

	fn := runtime.FuncForPC(pc)
	a = append([]interface{}{fn.Name(), filepath.Base(file) + "@" + strconv.Itoa(line) + ":"}, a...)

	l.Outputf(infoLog, "%s %s "+format, a...)
}

func (l *Logger) Warn(a ...interface{}) {
	pc, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		return
	}

	fn := runtime.FuncForPC(pc)
	a = append([]interface{}{fn.Name(), filepath.Base(file) + "@" + strconv.Itoa(line) + ":"}, a...)

	l.Output(warnLog, a...)
}

func (l *Logger) Warnf(format string, a ...interface{}) {
	pc, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		return
	}

	fn := runtime.FuncForPC(pc)
	a = append([]interface{}{fn.Name(), filepath.Base(file) + "@" + strconv.Itoa(line) + ":"}, a...)

	l.Outputf(warnLog, "%s %s "+format, a...)
}

func (l *Logger) Error(a ...interface{}) {
	if l.level <= errorLog {

		pc, file, line, ok := runtime.Caller(calldepth)
		if !ok {
			return
		}

		fn := runtime.FuncForPC(pc)
		a = append([]interface{}{fn.Name(), filepath.Base(file) + "@" + strconv.Itoa(line) + ":"}, a...)

		l.Output(errorLog, a...)
	}
}

func (l *Logger) Errorf(format string, a ...interface{}) {
	pc, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		return
	}

	fn := runtime.FuncForPC(pc)
	a = append([]interface{}{fn.Name(), filepath.Base(file) + "@" + strconv.Itoa(line) + ":"}, a...)

	l.Outputf(errorLog, "%s %s "+format, a...)
}

func (l *Logger) Fatal(a ...interface{}) {
	pc, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		return
	}

	fn := runtime.FuncForPC(pc)
	a = append([]interface{}{fn.Name(), filepath.Base(file) + "@" + strconv.Itoa(line) + ":"}, a...)

	l.Output(fatalLog, a...)
}

func (l *Logger) Fatalf(format string, a ...interface{}) {
	pc, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		return
	}

	fn := runtime.FuncForPC(pc)
	a = append([]interface{}{fn.Name(), filepath.Base(file) + "@" + strconv.Itoa(line) + ":"}, a...)

	l.Outputf(fatalLog, "%s %s "+format, a...)
}

func (l *Logger) Level() string {
	return levelName(l.level)
}

func (l *Logger) SetLevel(level uint8) {
	l.level = level
}