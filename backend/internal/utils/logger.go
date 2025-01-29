package utils

import (
	"log"
	"runtime"
)

type Logger struct {
	*log.Logger
}

// NewLogger 创建一个新的 Logger 实例
func NewLogger() *Logger {
	return &Logger{log.Default()}
}

// Println 输出日志，包含文件名和行号
func (l *Logger) Println(v ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		l.Logger.SetPrefix(file + ":" + string(line) + " ")
	}
	l.Logger.Println(v...)
}

// Printf 格式化输出日志，包含文件名和行号
func (l *Logger) Printf(format string, v ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		l.Logger.SetPrefix(file + ":" + string(line) + " ")
	}
	l.Logger.Printf(format, v...)
}
