// Goroutine(4)
package main

import (
	"fmt"
	_ "fmt"
	"runtime"
	"time"
	_ "time"
)

func main() {
	// Goroutine
	// Closure 사용 예제
	// 예제 1.
	runtime.GOMAXPROCS(2) // 사용하는 CPU CORE 개수 설정

	s := "Goroutine Closure : "

	for i := 0; i < 1000; i++ {
		go func(n int) {
			fmt.Println(s, n, " - ", time.Now())
		}(i)
		// 반복문 클로저는 일반적으로 즉시 실행 But 고루틴 클로저는 가장 나중에 실행 (반복문 종료 후)
	}

	time.Sleep(5 * time.Second)

}
