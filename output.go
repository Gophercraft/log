package log

type Output interface {
	Begin(lgr *Logger) error
	End() error
	AddLine(l *Line) error
	Include(cat Category) error
	Exclude(cat Category) error
}

type Console interface {
	Output
	AddProgressBar(pb *ProgressBar) error
	RemoveProgressBar(pb *ProgressBar) error
}
