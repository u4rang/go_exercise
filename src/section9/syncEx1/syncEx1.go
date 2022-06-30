// Synchronization 심화(1)
package main

import (
	"fmt"
	_ "fmt"
	"runtime"
	_ "runtime"
	"sync"
	_ "sync"
	"time"
	_ "time"
)

func onceTest() {
	// 이 부분에 한번 실행 할 코드 작성
	fmt.Println("once Test Execute!")
}

func main() {
	// Synchronization
	// Once : 한 번만 실행(주로 초기화에서 사용)
	// Do 로 실행

	runtime.GOMAXPROCS(runtime.NumCPU())

	once := new(sync.Once) // Once 객체 생성

	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println("goroutine : ", n)
			once.Do(onceTest) // golang Engine 에서 관리하며 1번만 실행하도록 처리
		}(i)
	}

	time.Sleep(2 * time.Second)
}
