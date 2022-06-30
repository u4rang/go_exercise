// Defer(3)
package main

import (
	"fmt"
)

// First In Last Out
func stack() {
	for i := 1; i <= 10; i++ {
		defer fmt.Println("stack : ", i)
	}
}

func main() {
	// Defer

	// 예제 1.
	stack()
}
