//go:build windows
// +build windows

package log

import "golang.org/x/sys/windows"

func eraseLine() {
	// get info about our console window
	var info windows.ConsoleScreenBufferInfo
	hConsole, err := windows.GetStdHandle(windows.STD_OUTPUT_HANDLE)
	if err != nil {
		panic(err)
	}
	windows.GetConsoleScreenBufferInfo(hConsole, &info)

	goBack(hConsole, &info)

	// write empty line to "clear" it of characters
	width := int(info.Window.Right - info.Window.Left + 1)
	buf := make([]uint16, width)
	for i := range buf {
		buf[i] = uint16(' ')
	}
	windows.WriteConsole(hConsole, &buf[0], uint32(width), nil, nil)

	goBack(hConsole, &info)
}

func goBack(hConsole windows.Handle, info *windows.ConsoleScreenBufferInfo) {
	position := info.CursorPosition

	position.Y -= 1
	position.X = 0

	windows.SetConsoleCursorPosition(hConsole, position)
}
