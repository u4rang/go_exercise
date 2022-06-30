// channel(6)
package main

import (
	"fmt"
	_ "fmt"
	_ "time"
)

func main() {
	// Channel
	// Close : 채널 닫기

	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- 7777
		}
	}()

	val1, ok1 := <-ch
	fmt.Println("val1 : ", val1, " / ok1 : ", ok1)
	val2, ok2 := <-ch
	fmt.Println("val2 : ", val2, " / ok2 : ", ok2)
	val3, ok3 := <-ch
	fmt.Println("val3 : ", val3, " / ok3 : ", ok3)

	close(ch) // Channel close

	val4, ok4 := <-ch
	fmt.Println("val4 : ", val4, " / ok4 : ", ok4)

}
