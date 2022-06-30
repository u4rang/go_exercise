// Synchronization 심화(3)
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

func main() {
	// Synchronization - Atomic, 원자성
	// All or Nothing (ex. 계좌이체)
	// - 기능적으로 분할 불가능한 완전 보증된 일련의 작업, 모두 성공 또는 모두 실패
	// - 모든 조작이 완료 될 때까지 다른 프로세스 개입 불가
	// sync/atomic 에서 원자적 연산자 제공
	// https://golang.org/pkg/sync/atomic 에서 계열 확인 가능
	// 주로 공용 변수에 관한 계산 사용

	// 원자성 사용 안하는 예제
	runtime.GOMAXPROCS(runtime.NumCPU())

	var cnt int64 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(n int) {
			cnt += 1
			wg.Done()
		}(i)
	}

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(n int) {
			cnt -= 1
			wg.Done()
		}(i)

	}

	// Add(7000) == Done(7000) 횟수가 같아야 한다.
	wg.Wait()

	fmt.Println("WaitGroup End! cnt :", cnt)

}
