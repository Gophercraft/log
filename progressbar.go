package log

import (
	"errors"
	"sync"
	"time"
)

var ErrNoConsole = errors.New("log: no console to use progress bar on")

type ProgressBar struct {
	Range ProgressRange
	Title string

	start time.Time

	protect sync.Mutex
	update  chan bool

	oscillationIndex int
	lastOscillation  time.Time

	l *Logger
}

func (pb *ProgressBar) requestUpdate() {
	if pb.update != nil {
		pb.update <- true
	}
}

func (pb *ProgressBar) SetFloat(f float64) {
	pb.protect.Lock()
	pb.Range.SetFloat(f)
	pb.protect.Unlock()

	pb.requestUpdate()
}

func (pb *ProgressBar) SetInt(i int64) {
	pb.protect.Lock()
	pb.Range.SetInt(i)
	pb.protect.Unlock()

	pb.requestUpdate()
}

type ProgressRange interface {
	MinMax() (string, string)
	PercentValue() float64
	Value() string
	SetFloat(f float64)
	SetInt(i int64)
}

func (l *Logger) StartProgressBar(pb *ProgressBar) error {
	pb.l = l
	pb.start = time.Now()

	for _, op := range l.Outputs {
		if cons, ok := op.(Console); ok {
			return cons.AddProgressBar(pb)
		}
	}

	return ErrNoConsole
}

func (l *Logger) RemoveProgressBar(pb *ProgressBar) error {
	for _, op := range l.Outputs {
		if cons, ok := op.(Console); ok {
			return cons.RemoveProgressBar(pb)
		}
	}

	return ErrNoConsole
}

func StartProgressBar(pb *ProgressBar) {
	if err := DefaultLogger.StartProgressBar(pb); err != nil {
		panic(err)
	}
}

func (pb *ProgressBar) Complete() error {
	return pb.l.RemoveProgressBar(pb)
}
