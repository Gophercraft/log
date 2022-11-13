package log

import "strconv"

type ProgressInt struct {
	mn, mx int64
	value  int64
}

func itos(i int64) string {
	return strconv.FormatInt(i, 10)
}

func (pi *ProgressInt) MinMax() (string, string) {
	return itos(pi.mn), itos(pi.mx)
}

func (pi *ProgressInt) PercentValue() float64 {
	progress := ((float64(pi.value) - float64(pi.mn)) * 100) / (float64(pi.mx) - float64(pi.mn))
	return progress
}

func (pi *ProgressInt) Value() string {
	return itos(pi.value)
}

func (pi *ProgressInt) SetFloat(f float64) {
	pi.value = int64(f)
}

func (pi *ProgressInt) SetInt(i int64) {
	pi.value = i
}

func NewIntProgressBar(title string, min, max int64) *ProgressBar {
	pb := new(ProgressBar)
	pb.Title = title
	pb.Range = &ProgressInt{
		mn: min,
		mx: max,
	}
	return pb
}
