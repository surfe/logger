package simple

import (
	"log"
)

type Wrapper struct {
	logger *log.Logger
}

func (w *Wrapper) Errorf(template string, err interface{}, args ...interface{}) {
	log.Printf(template, err, args)
}

func (w *Wrapper) Errorw(msg string, err interface{}, keysAndValues ...interface{}) {
	log.Printf(msg+"%v %v ", keysAndValues, err)
}

func (w *Wrapper) Error(err interface{}, args ...interface{}) {
	log.Printf("%v %v", err, args)
}

func (w *Wrapper) Infof(template string, args ...interface{}) {
	log.Printf(template, args)
}

func (w *Wrapper) Infow(msg string, keysAndValues ...interface{}) {
	log.Printf(msg+"%v ", keysAndValues)
}

func (w *Wrapper) Info(args ...interface{}) {
	log.Printf("%v ", args)
}

func (w *Wrapper) Debugf(template string, args ...interface{}) {
	log.Printf(template, args)
}

func (w *Wrapper) Debugw(msg string, keysAndValues ...interface{}) {
	log.Printf(msg+"%v ", keysAndValues)
}

func (w *Wrapper) Debug(args ...interface{}) {
	log.Printf("%v ", args)
}

func (w *Wrapper) Sync() {
	// Nothing to sync
}
