package log

import (
	"fmt"
	"os"
	"path/filepath"
)

type DirOutput struct {
	Dir string

	cancel chan bool
	msg    chan *Line

	exclude map[Category]bool
	files   map[Category]*os.File
}

func (do *DirOutput) print(msg *Line) {
	if do.exclude[msg.Category] {
		return
	}

	var file *os.File
	var ok bool
	var err error
	file, ok = do.files[msg.Category]
	if !ok {
		file, err = os.OpenFile(filepath.Join(do.Dir, fmt.Sprintf("%s.log.txt", msg.Category)), os.O_CREATE|os.O_APPEND, 0700)
		if err != nil {
			return
		}
		do.files[msg.Category] = file
	}

	fmt.Fprintf(file, "%s] %s\n", printTime(msg.Time), msg.Text)
}

func (do *DirOutput) closeAll() {
	for _, file := range do.files {
		file.Close()
	}

	do.files = nil
}

func (do *DirOutput) handle() {
	for {
		select {
		case <-do.cancel:
			do.closeAll()
			return
		case msg := <-do.msg:
			do.print(msg)
		}
	}
}

func (do *DirOutput) Begin(l *Logger) error {
	if err := os.MkdirAll(do.Dir, 0700); err != nil {
		return err
	}

	do.exclude = make(map[Category]bool)
	do.files = make(map[Category]*os.File)
	do.cancel = make(chan bool)
	do.msg = make(chan *Line)
	go do.handle()
	return nil
}

func (do *DirOutput) End() error {
	do.cancel <- true
	close(do.cancel)
	close(do.msg)
	return nil
}

func (do *DirOutput) Exclude(cat Category) error {
	do.exclude[cat] = true
	return nil
}

func (do *DirOutput) Include(cat Category) error {
	delete(do.exclude, cat)
	return nil
}

func (do *DirOutput) AddLine(ln *Line) error {
	return nil
}
