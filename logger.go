package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
)

type Logger struct {
	prefix   string
	filename string
}

func New(filename string, prefix string) *Logger {
	return &Logger{
		filename: filename,
		prefix:   prefix,
	}
}

func (log Logger) Write(msg string, TYPE string) {
	now := time.Now().UTC().Format("2006-01-02 UTC 15:04:05.000")
	out := fmt.Sprintf("[%s] %s [%s] %s\n", log.prefix, now, TYPE, msg)
	file, _ := os.OpenFile(log.filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer file.Close()
	file.Write([]byte(out))
}

func (log Logger) D(msg string) {
	log.Write(msg, "DEBUG")
}

func (log Logger) I(msg string) {
	log.Write(msg, "INFOM")
}

func (log Logger) E(msg string) {
	log.Write(msg, "ERROR")
}

func (log Logger) Err(err error) {
	log.Write(fmt.Sprintf("%+v", err), "ERROR")
}
