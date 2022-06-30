// Goroutine(2)
package main

import (
	"fmt"
	_ "fmt"
	"time"
	_ "time"
)

func exe(name string) {
	fmt.Println(name, "start :", time.Now())

	for i := 0; i < 100000; i++ {
		fmt.Println(name, ">>>>", i)
	}

	fmt.Println(name, "end :", time.Now())
}

func main() {
	// Goroutine
	exe("t1")

	// 예제 1.
	fmt.Println("Main Routine Start :", time.Now())
	go exe("t2")
	go exe("t3")
	time.Sleep(4 * time.Second) // time.Second, time.Minute, time.Hour, time.Milisecond

	fmt.Println("Main Routine End :", time.Now())
}
