//go:build !windows

package log

import "fmt"

func eraseLine() {
	fmt.Print("\033[1A\033[K")
}
