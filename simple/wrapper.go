package simple

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/surfe/logger/logi"
)

type Logger struct {
	toAppend string
}

func (w *Logger) Ctx(ctx context.Context) logi.Logger {
	return w
}

func (w *Logger) With(args ...any) logi.Logger {
	var sb strings.Builder

	if len(w.toAppend) > 0 {
		sb.WriteString(w.toAppend)
		sb.WriteString(";")
	}

	for i, a := range args {
		switch v := a.(type) {
		case int:
			sb.WriteString(strconv.Itoa(v))
		case float64:
			sb.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
		case string:
			sb.WriteString(v)
		case bool:
			sb.WriteString(strconv.FormatBool(v))
		default:
			sb.WriteString(fmt.Sprintf("%v", v))
		}

		if i%2 == 0 {
			sb.WriteString(": ")
		} else {
			sb.WriteString("; ")
		}
	}

	return &Logger{toAppend: sb.String()}
}

func (w *Logger) Err(err error) logi.Logger {
	return w.With("error", err)
}

func (w *Logger) Errorf(template string, args ...any) {
	w.printf(template, args...)
}

func (w *Logger) Error(args ...any) {
	w.println(args...)
}

func (w *Logger) Warnf(template string, args ...any) {
	w.printf(template, args...)
}

func (w *Logger) Warn(args ...any) {
	w.println(args...)
}

func (w *Logger) Infof(template string, args ...any) {
	w.printf(template, args...)
}

func (w *Logger) Info(args ...any) {
	w.println(args...)
}

func (w *Logger) Debugf(template string, args ...any) {
	w.printf(template, args...)
}

func (w *Logger) Debug(args ...any) {
	w.println(args...)
}

func (w *Logger) Fatalf(template string, args ...any) {
	w.printf(template, args...)
	os.Exit(1)
}

func (w *Logger) Fatal(args ...any) {
	w.println(args...)
	os.Exit(1)
}

func (w *Logger) Sync() {
	// Nothing to sync
}

func (w *Logger) println(args ...any) {
	log.Println(w.toAppend, fmt.Sprint(args...))
}

func (w *Logger) printf(template string, args ...any) {
	log.Printf(w.toAppend+template, args...)
}
