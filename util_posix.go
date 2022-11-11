//go:build !windows

package log

func eraseLine() {
	fmt.Print("\033[1A\033[K")
}
