// Defer(2)
package main

import (
	"fmt"
)

func sayHello(m string) {
	defer func() {
		fmt.Println("Hello!", m)
	}()

	func() {
		fmt.Println("Hi!", m)
	}()
}

func main() {
	// Defer

	// 예제 1.
	sayHello("Golang")
}
