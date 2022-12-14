package simple

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/surfe/logger/logi"
)

type Logger struct {
	toAppend string
}

func (w *Logger) With(ctx context.Context, args ...any) logi.Logger {
	toAppend := w.toAppend + " " + fmt.Sprint(args...)

	return &Logger{toAppend: toAppend}
}

func (w *Logger) Err(err error) logi.Logger {
	return w.With(context.TODO(), "error", err)
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
	print := w.toAppend + "; " + fmt.Sprint(args...)
	log.Println(print)
}

func (w *Logger) printf(template string, args ...any) {
	log.Printf(w.toAppend+"; "+template, args...)
}
