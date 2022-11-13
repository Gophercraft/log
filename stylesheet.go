package log

import (
	"time"

	"github.com/fatih/color"
)

type Color color.Attribute

var (
	Cyan   Color = Color(color.FgCyan)
	White  Color = Color(color.FgHiWhite)
	Yellow Color = Color(color.FgYellow)
	Red    Color = Color(color.FgHiRed)
)

type StyleSheet struct {
	Colors map[Category]Color

	TimeColor Color // color of timestamp

	// [=======> ]

	BarCaseLeft  rune   // '['
	BarCaseRight rune   // ']'
	BarLiquid    rune   // '='
	BarVoid      rune   // ' '
	BarHead      []rune // { '>' }

	HeadOscillation time.Duration // time needed to increment the barhead index
}

var DefaultStyleSheet = StyleSheet{
	Colors: map[Category]Color{
		"debug": White,
		"warn":  Yellow,
		"error": Red,
	},

	TimeColor: Cyan,

	BarCaseLeft:  '[',
	BarCaseRight: ']',

	BarLiquid: '=',
	BarVoid:   ' ',

	BarHead: []rune{'|', '╱', '─', '╲'},

	HeadOscillation: 300 * time.Millisecond,
}
