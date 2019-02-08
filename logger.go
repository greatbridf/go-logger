package logger

import (
	"fmt"
	"io"
	"time"
)

type Logger struct {
	prefix string
	out    io.Writer
}

func New(out io.Writer, prefix string) *Logger {
	return &Logger{
		out:    out,
		prefix: prefix,
	}
}

func (log Logger) Write(msg string, TYPE string) {
	now := time.Now().UTC().Format("2006-01-02 UTC 15:04:05")
	out := fmt.Sprintf("[%s] %s [%s] %s\n", log.prefix, now, TYPE, msg)
	log.out.Write([]byte(out))
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
