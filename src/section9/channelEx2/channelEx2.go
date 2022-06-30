// channel 심화(2)

package main

import (
	"fmt"
	_ "fmt"
	_ "time"
)

func sum(cnt int) <-chan int {
	sum := 0
	tot := make(chan int)
	go func() {
		for i := 0; i < cnt; i++ {
			sum += i
		}

		tot <- sum
	}()

	return tot
}

func main() {
	// Channel
	// 채널 또한 함수의 반환 값으로 사용 가능

	// 예제 1.
	c := sum(100)

	fmt.Println("sum(100) :", <-c)

}
