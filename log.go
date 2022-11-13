package log

import (
	"time"
)

type Category string

type Line struct {
	Time     time.Time
	Category Category
	Text     string
	Fatal    bool
}

type Logger struct {
	StyleSheet StyleSheet
	Outputs    []Output
}

var DefaultLogger Logger

func init() {
	DefaultLogger.StyleSheet = DefaultStyleSheet
	DefaultLogger.AddOutput(&ConsoleOutput{})
}

func (l *Logger) AddOutput(op Output) error {
	l.Outputs = append(l.Outputs, op)
	if err := op.Begin(l); err != nil {
		return err
	}

	return nil
}

func (l *Logger) LogLine(ln *Line) {
	for _, output := range l.Outputs {
		output.AddLine(ln)
	}
}

func (l *Logger) GetConsole() *ConsoleOutput {
	for _, op := range l.Outputs {
		switch o := op.(type) {
		case *ConsoleOutput:
			return o
		}
	}
	return nil
}

func (l *Logger) GetDir() *DirOutput {
	for _, op := range l.Outputs {
		switch o := op.(type) {
		case *DirOutput:
			return o
		}
	}
	return nil
}

func (l *Logger) ExcludeConsole(cat Category) {
	d := l.GetConsole()
	if d != nil {
		d.Exclude(cat)
	}
}

func (l *Logger) ExcludeDir(cat Category) {
	d := l.GetDir()
	if d != nil {
		d.Exclude(cat)
	}
}

func (l *Logger) IncludeConsole(cat Category) {
	d := l.GetConsole()
	if d != nil {
		d.Include(cat)
	}
}

func (l *Logger) IncludeDir(cat Category) {
	d := l.GetDir()
	if d != nil {
		d.Include(cat)
	}
}
