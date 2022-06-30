// Goroutine(3)
package main

import (
	"fmt"
	_ "fmt"
	"math/rand"
	_ "math/rand"
	"runtime"
	_ "runtime"
	"time"
	_ "time"
)

func exe(name int) {
	r := rand.Intn(100) // 100 이하의 랜덤함 수
	fmt.Println(name, "start :", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, ">>>>", r, i)
	}
	fmt.Println(name, "end :", time.Now())
}

func main() {
	// Goroutine
	// - Multi Core(다중 CPU) 최대한 활용

	runtime.GOMAXPROCS(runtime.NumCPU())                       // 현 시스템의 CPU 코어 개수 반환 후 설정
	fmt.Println("Current System CPU :", runtime.GOMAXPROCS(0)) // 설정 값 출력

	// 예제 1.
	fmt.Println("Main Routine Start :", time.Now())

	// Goroutine 100개 생성
	for i := 0; i < 1000; i++ {
		go exe(i)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Main Routine End :", time.Now())
}
