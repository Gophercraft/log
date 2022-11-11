package log

import (
	"fmt"
	"time"
)

func (l *Logger) Println(args ...any) {
	str := fmt.Sprint(args...)
	l.LogLine(&Line{time.Now(), "debug", str})
}

func (l *Logger) Printf(format string, args ...any) {
	l.LogLine(&Line{time.Now(), "debug", fmt.Sprintf(format, args...)})
}

func (l *Logger) CatPrintf(cat Category, format string, args ...any) {
	l.LogLine(&Line{time.Now(), cat, fmt.Sprintf(format, args...)})
}

// shorthand

func Println(args ...any) {
	DefaultLogger.Println(args...)
}

func Printf(format string, args ...any) {
	DefaultLogger.Printf(format, args...)
}
