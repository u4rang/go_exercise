// channel(2)
package main

import (
	"fmt"
	_ "fmt"
	_ "time"
)

func rangeSum(rg int, c chan int) {
	sum := 0

	for i := 0; i < rg; i++ {
		sum += i
	}

	c <- sum
}

func main() {
	// Channel

	c := make(chan int)

	go rangeSum(1000000, c)
	go rangeSum(7000, c)
	go rangeSum(5000, c)

	// 순서대로 대에티 수신(동기) : 채널에서 값 수신 완료 될 때까지 대기
	result1 := <-c
	result2 := <-c
	result3 := <-c

	fmt.Println("result 1:", result1)
	fmt.Println("result 2:", result2)
	fmt.Println("result 3:", result3)
}
