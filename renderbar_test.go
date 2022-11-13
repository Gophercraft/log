package log

import "testing"

func TestRenderbar(t *testing.T) {
	sheet := DefaultStyleSheet
	sheet.BarHead = []rune{'>'}

	for _, test := range []struct {
		Width    int
		Percent  float64
		Expected string
	}{
		{
			Width:    10,
			Percent:  50,
			Expected: "[===>    ]",
		},
		{
			Width:    10,
			Percent:  100,
			Expected: "[=======>]",
		},
		{
			Width:    4,
			Percent:  21.3,
			Expected: "[> ]",
		},
	} {
		bar := renderBarInWidth(&sheet, 0, test.Width, test.Percent)
		if bar != test.Expected {
			t.Fatalf("Failure, expected '%s', got '%s'", test.Expected, bar)
		}
	}

}
