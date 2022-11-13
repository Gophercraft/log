package log

func renderBarInWidth(sheet *StyleSheet, oscillationIndex int, availableChars int, percent float64) string {
	bar := make([]rune, availableChars)

	barInsideLength := availableChars - 2
	barInsideProgress := int(float64(barInsideLength) * (percent / 100))

	// encase characters
	bar[0] = sheet.BarCaseLeft
	bar[availableChars-1] = sheet.BarCaseRight

	for i := 1; i < barInsideProgress; i++ {
		bar[i] = sheet.BarLiquid
	}

	if barInsideProgress == 0 {
		barInsideProgress = 1
	}

	bar[barInsideProgress] = sheet.BarHead[oscillationIndex]

	for i := barInsideProgress + 1; i < availableChars-1; i++ {
		bar[i] = sheet.BarVoid
	}

	return string(bar)
}
