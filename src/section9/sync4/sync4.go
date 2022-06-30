// Synchronization(4)
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
	// Synchronization
	// Mutex 종류
	// 1. RWMutex : 쓰기 Lock -> 쓰기 시도 중에는 다른곳에서 이전 값을 읽으면 X / 읽기 락, 쓰기 락 전부 방어(방지)
	// 2. RMutex : 읽기 Lock -> 읽기 시도 중 값 변경 방지
	//                         즉, 쓰기 락 방어(방지)

	// 시스템 전체 CPU 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0
	mutex := new(sync.RWMutex) // 쓰기 Lock 선언
	// var mutex = new(sync.RWMutex)

	go func() {
		for i := 1; i <= 10; i++ {
			// 쓰기 mutex 잠금
			mutex.Lock()

			data += 1
			fmt.Println("Write1 :", data)
			time.Sleep(200 * time.Millisecond)

			// 쓰기 mutex 해제
			mutex.Unlock()
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			// 쓰기 mutex 잠금
			mutex.Lock()

			data += 1
			fmt.Println("Write2 :", data)
			time.Sleep(200 * time.Millisecond)

			// 쓰기 mutex 해제
			mutex.Unlock()
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			// 읽기 mutex 잠금
			mutex.RLock()

			fmt.Println("Read1 :", data)
			time.Sleep(200 * time.Microsecond)

			// 읽기 mutex 해제
			mutex.RUnlock()
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			// 읽기 mutex 잠금
			mutex.RLock()

			fmt.Println("Read2 :", data)
			time.Sleep(1 * time.Second)

			// 읽기 mutex 해제
			mutex.RUnlock()
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			// 읽기 mutex 잠금
			mutex.RLock()

			fmt.Println("Read3 :", data)
			time.Sleep(1 * time.Second)

			// 읽기 mutex 해제
			mutex.RUnlock()

		}
	}()

	time.Sleep(15 * time.Second)
}
