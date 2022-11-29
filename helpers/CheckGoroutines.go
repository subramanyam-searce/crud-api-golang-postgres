package helpers

import (
	"fmt"
	"runtime"
)

func CheckGoroutines() {
	goroutines := runtime.NumGoroutine()

	if goroutines > 5 {
		fmt.Println("Multiple Goroutines Detected:", goroutines)
	}
}
