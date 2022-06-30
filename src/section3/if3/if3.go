// IF문(3)
package main

import "fmt"

func main() {
	i := 100

	// if -else if 예제(1)
	if i >= 120 {
		fmt.Println("Case 1")
	} else if i >= 100 && i < 120 {
		fmt.Println("Case 2")
	} else if i < 100 && i >= 50 {
		fmt.Println("Case 3")
	} else {
		fmt.Println("Case 4")
	}
}
