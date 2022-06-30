// channel 심화(1)

package main

import (
	"fmt"
	_ "fmt"
	"time"
	_ "time"
)

func main() {
	// Channel - 셀랙트 구문
	// 채널 Select 구문 -> 채널에 값이 수신되면 해당 Case 문 실행
	// 일회성 구문이므로, For(반복문) 안에서 수행
	// Default 구문 처리 주의 <- 사용하지 않기

	// 예제 1.
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- 77
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "Hi! golang~"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case num := <-ch1:
				fmt.Println("ch1 : ", num)
			case str := <-ch2:
				fmt.Println("ch2 : ", str)
				// default:
				// 	fmt.Println("Default test")
			}
		}
	}()

	time.Sleep(7 * time.Second)
}
