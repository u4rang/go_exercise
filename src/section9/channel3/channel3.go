// channel(3)
package main

import (
	"fmt"
	_ "fmt"
	"time"
	_ "time"
)

func main() {
	// Channel

	// 예제 1 - 동기 : 버퍼 미사용
	ch := make(chan bool)
	cnt := 6

	// 익명함수
	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true // 1. 송신하고 대기
			fmt.Println("Go :", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch // 2. 수신 후 익명함수 진행
		fmt.Println("Main :", i)
	}
}
