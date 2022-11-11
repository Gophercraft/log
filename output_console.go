package log

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/nathan-fiscaletti/consolesize-go"
)

type ConsoleOutput struct {
	l *Logger
	// causes the progress bars to refresh
	ticker *time.Ticker
	// causes the output handler to shutdown
	cancel chan bool
	// causes the progress bars to refresh
	refresh chan bool
	// causes line buffers to appear in-order
	msg chan *Line
	// allows progressbars to be added
	addPbChan chan *ProgressBar
	// allows progressbars to be removed
	rmPbChan chan *ProgressBar
	// list of active progressbars
	progressbars []*ProgressBar
	// the number of progressbars currently displayed in console (alternatively: how many lines to remove!)
	displayedProgressBarCount int
}

func (w *ConsoleOutput) eraseLines(lines int) {
	for i := 0; i < lines; i++ {
		eraseLine()
	}
}

func (w *ConsoleOutput) getConsoleWidth() int {
	cols, _ := consolesize.GetConsoleSize()
	return cols
}

func formatDur(d time.Duration) string {
	if d < 1*time.Minute {
		return fmt.Sprintf("%ds", d/time.Second)
	}

	if d < 1*time.Hour {
		return fmt.Sprintf("%dm%ds", d/time.Minute, d/time.Second)
	}

	return d.String()
}

func (w *ConsoleOutput) renderProgressBar(pb *ProgressBar) {
	width := w.getConsoleWidth()

	titleCharacters := []rune(pb.Title)

	pbAvailableChars := width
	pbAvailableChars -= len(titleCharacters)
	pbAvailableChars -= 1 // space

	mn, mx := pb.Range.MinMax()

	progressText := fmt.Sprintf(" %s %s/%s (%.2f%%) %s", pb.Title, mn, mx, pb.Range.PercentValue(), formatDur(time.Since(pb.start)))

	pbAvailableChars -= len([]rune(progressText))

	bar := renderBarInWidth(&w.l.StyleSheet, pb.oscillationIndex, pbAvailableChars, pb.Range.PercentValue())

	if time.Since(pb.lastOscillation) > w.l.StyleSheet.HeadOscillation {
		pb.oscillationIndex++
		if pb.oscillationIndex >= len(w.l.StyleSheet.BarHead) {
			pb.oscillationIndex = 0
		}
		pb.lastOscillation = time.Now()
	}

	fmt.Printf("%s%s\n", bar, progressText)
}

func (w *ConsoleOutput) eraseProgressBars() {
	linesToErase := w.displayedProgressBarCount
	w.eraseLines(linesToErase)
	w.displayedProgressBarCount = 0
}

func (w *ConsoleOutput) renderProgressBars() {
	for _, pb := range w.progressbars {
		w.renderProgressBar(pb)
	}
	w.displayedProgressBarCount = len(w.progressbars)
}

func (w *ConsoleOutput) refreshProgressBars() {
	w.eraseProgressBars()
	w.renderProgressBars()
}

func (w *ConsoleOutput) print(ln *Line) {
	w.eraseProgressBars()

	color.Set(color.Attribute(w.l.StyleSheet.TimeColor))

	fmt.Printf("[%s] ", printTime(ln.Time))

	color.Unset()

	cl, setColor := w.l.StyleSheet.Colors[ln.Category]
	if setColor {
		color.Set(color.Attribute(cl))
	}

	fmt.Printf("[%s] ", ln.Category)

	if setColor {
		color.Unset()
	}

	fmt.Printf("%s\r\n", ln.Text)

	w.refreshProgressBars()
}

func (w *ConsoleOutput) appendProgressBar(pb *ProgressBar) {
	w.progressbars = append(w.progressbars, pb)
}

func (w *ConsoleOutput) removeProgressBar(pb *ProgressBar) {
	index := -1

	for i, curPb := range w.progressbars {
		if curPb == pb {
			index = i
			break
		}
	}

	if index >= 0 {
		w.progressbars = append(w.progressbars[:index], w.progressbars[index+1:]...)
	}
}

func (w *ConsoleOutput) begin() {
	for {
		select {
		case <-w.cancel:
			return
		case pb := <-w.addPbChan:
			w.appendProgressBar(pb)
		case pb := <-w.rmPbChan:
			w.removeProgressBar(pb)
		case <-w.refresh:
			w.refreshProgressBars()
		case msg := <-w.msg:
			w.print(msg)
		case <-w.ticker.C:
			w.refreshProgressBars()
		}
	}
}

func (w *ConsoleOutput) Begin(l *Logger) error {
	w.l = l
	w.cancel = make(chan bool)
	w.msg = make(chan *Line)
	w.refresh = make(chan bool)
	w.addPbChan = make(chan *ProgressBar)
	w.rmPbChan = make(chan *ProgressBar)
	w.ticker = time.NewTicker(200 * time.Millisecond)
	go w.begin()
	return nil
}

func (w *ConsoleOutput) End() error {
	w.cancel <- true
	close(w.cancel)
	close(w.refresh)
	close(w.addPbChan)
	close(w.rmPbChan)
	w.ticker.Stop()
	return nil
}

func (w *ConsoleOutput) AddLine(l *Line) error {
	w.msg <- l
	return nil
}

func (w *ConsoleOutput) AddProgressBar(pb *ProgressBar) error {
	pb.update = w.refresh
	w.addPbChan <- pb
	return nil
}

func (w *ConsoleOutput) RemoveProgressBar(pb *ProgressBar) error {
	w.rmPbChan <- pb
	return nil
}

func (w *ConsoleOutput) Exclude(cat Category) error {
	return nil
}

func (w *ConsoleOutput) Include(cat Category) error {
	return nil
}
