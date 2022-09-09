package logging

import (
	"fmt"
	"gin_log/pkg/cron"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int
var (
	F *os.File

	DefaultPrefix = ""
	DefaultCallerDepth = 2

	logger *log.Logger
	logPrefix = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)
const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
	go cron.ScheduleTask(updateF)
}
func updateF() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}
func Debug(v ...any) {
	setPrefix(DEBUG)
	logger.Println(v)
}
func Info(v ...any) {
	setPrefix(INFO)
	logger.Println(v)
}
func Warn(v ...any) {
	setPrefix(WARN)
	logger.Println(v)
}
func Error(v ...any) {
	setPrefix(ERROR)
	logger.Println(v)
}
func Fatal(v ...any) {
	setPrefix(FATAL)
	logger.Fatal(v)
}
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	}else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
