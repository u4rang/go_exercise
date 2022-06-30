// Synchronization(5)
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

func main() {
	// Synchronization - 실행 흐름 제어
	// goroutine 동기화 객체
	// 동기화 상태(조건) 메소드 사용
	// Wait, notify, notifyAll : 기타 언어
	// Wait, Signal, Broadcast

	// 시스템 전체 CPU 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex) // 범용 사용
	var condition = sync.NewCond(mutex)

	c := make(chan int, 5) // 비동기 버퍼 채널

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			fmt.Println("goroutine Start : ", n)
			c <- 777
			fmt.Println("goroutine waiting : ", n)
			condition.Wait() // goroutine waiting(대기, 멈춤)
			fmt.Println("goroutine End : ", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		<-c
		// fmt.Println("Received : ", <-c)
	}

	/*
		//  한 개씩 깨우기
			for i := 0; i < 5; i++ {
				mutex.Lock()

				fmt.Println("Wake goroutine(Signal) : ", i)
				condition.Signal() // 한 개씩 깨운다.

				mutex.Unlock()
			}
	*/

	// 한 번에 깨우기
	mutex.Lock()
	fmt.Println("Wake goroutine(Broadcast)")
	condition.Broadcast()
	mutex.Unlock()

	time.Sleep(2 * time.Second)

}
