// Synchronization 심화(2)
package main

import (
	"fmt"
	_ "fmt"
	_ "runtime"
	"sync"
	_ "sync"
	_ "time"
)

func main() {
	// Synchronization - 대기 그룹
	// 실행 흐름 변경 (고루틴이 종료 될 때 까지 대기 가능)
	// WaitGroup : Add(고루틴 추가), Done(작업 종료 알림), Wait(고루틴 종료시 까지 대기)
	// Add 로 추가 된 고루틴 개수와 Done 으로 종료되는 알림 횟수 같아야 정확하게 동작한다. (ADD == Done)

	wg := new(sync.WaitGroup)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println("WaitGroup: ", n)
			wg.Done()
		}(i)
	}

	// Add == Done 횟수가 같아야 한다.
	wg.Wait()

	fmt.Println("WaitGroup End!")

}
