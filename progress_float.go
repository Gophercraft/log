package log

import "strconv"

type ProgressFloat struct {
	mn, mx float64
	value  float64
}

func ftos(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64)
}

func (pf *ProgressFloat) MinMax() (string, string) {
	return ftos(pf.mn), ftos(pf.mx)
}

func (pf *ProgressFloat) PercentValue() float64 {
	progress := ((pf.value - pf.mn) * 100) / (pf.mx - pf.mn)
	return progress
}

func (pf *ProgressFloat) Value() string {
	return ftos(pf.value)
}

func (pf *ProgressFloat) SetFloat(f float64) {
	pf.value = f
}

func (pf *ProgressFloat) SetInt(i int64) {
	pf.value = float64(i)
}

func NewFloatProgressBar(title string, min, max float64) *ProgressBar {
	pb := new(ProgressBar)
	pb.Title = title
	pb.Range = &ProgressFloat{
		mn: min,
		mx: max,
	}
	return pb
}
