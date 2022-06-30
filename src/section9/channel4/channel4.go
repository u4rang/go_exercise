// channel(1)
package main

import (
	"fmt"
	_ "fmt"
	"runtime"
	_ "runtime"
	_ "time"
)

func main() {
	// Channel - Buffer
	// Buffer
	// - 발신 : 가득차면 대기, 비어있으면 작동
	// - 수신 : 비어있으면 대기, 가득차있으면 작동

	// 예제 1 - 비동기 : 버퍼 사용

	runtime.GOMAXPROCS(1)
	ch := make(chan bool, 5) // Channel에 버퍼의 용량 설정
	cnt := 12

	// 익명함수
	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go :", i)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main :", i)
	}
}
