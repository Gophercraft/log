package log

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func (l *Logger) Println(args ...any) {
	str := fmt.Sprint(args...)
	l.LogLine(&Line{time.Now(), "debug", str, false})
}

func (l *Logger) Printf(format string, args ...any) {
	l.LogLine(&Line{time.Now(), "debug", fmt.Sprintf(format, args...), false})
}

func (l *Logger) CatPrintf(cat Category, format string, args ...any) {
	l.LogLine(&Line{time.Now(), cat, fmt.Sprintf(format, args...), false})
}

// shorthand

// Prints 1+ items to default logger
func Println(args ...any) {
	DefaultLogger.Println(args...)
}

// Prints a formatted string to default logger
func Printf(format string, args ...any) {
	DefaultLogger.Printf(format, args...)
}

// Prints a warning to the default logger
func Warn(args ...any) {
	str := fmt.Sprint(args...)
	DefaultLogger.LogLine(&Line{time.Now(), "warn", str, false})
}

// Prints 1+ error messages to the default logger
func Err(args ...any) {
	str := fmt.Sprint(args...)
	DefaultLogger.LogLine(&Line{time.Now(), "error", str, false})
}

func death() {
	// Bide your time and wait for the end
	wait := make(chan bool)
	<-wait
}

// Crash the program, providing 1+ error messages
func Fatal(args ...any) {
	str := fmt.Sprint(append([]any{"Fatal: "}, args...)...)
	DefaultLogger.LogLine(&Line{time.Now(), "error", str, true})
	death()
}

func dump(object any) string {
	return spew.Sdump(object)
}

func NewDumpLine(name string, cat Category, object any) *Line {
	return &Line{
		time.Now(),
		cat,
		fmt.Sprintf("%s = %s", name, dump(object)),
		false,
	}
}

// Show a formatted representation of an object to the default logger
func Dump(name string, object any) {
	ln := NewDumpLine(name, "debug", object)
	DefaultLogger.LogLine(ln)
}

// Show a formatted object representation before crashing the program
func FatalDump(name string, object any) {
	ln := NewDumpLine(name, "error", object)
	ln.Fatal = true
	DefaultLogger.LogLine(ln)
	death()
}
