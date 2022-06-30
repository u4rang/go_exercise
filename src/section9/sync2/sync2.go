// Synchronization(2)
package main

import (
	"fmt"
	_ "fmt"
	"runtime"
	_ "runtime"
	"sync"
	_ "sync"
	_ "time"
)

type count struct {
	num   int
	mutex sync.Mutex
}

func (c *count) increment() {
	c.mutex.Lock() // 공유 데이터 작업 전, 뮤텍스로 문 잠기
	c.num += 1
	c.mutex.Unlock() // 작업 완료 후, 뮤덱스로 문 열어놓기
}

func (c *count) result() {
	fmt.Println("result :", c.num)
}

func main() {
	// Synchronization - 동기화 처리 예제
	// goroutine synchronization
	// - 실행 흐름 제어 및 변수 동기화 가능
	// - 공유 데이터 보호가 가장 중요
	// - mutex : 여러 고루틴에서 작업하는 공유데이터 보호
	//           sync.Mutex 선언 후 Lock, Unlock 사용

	// 동기화 사용하는  예제
	// 시스템의 전체 CPU 사용
	runtime.GOMAXPROCS(runtime.NumCPU()) // 하이퍼쓰레딩 지원하는 코어 모두 사용 설정

	c := count{num: 0}
	done := make(chan bool)

	for i := 1; i <= 10000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched() // CPU 양보
		}()
	}

	for i := 1; i <= 10000; i++ {
		<-done

	}

	c.result()
}
