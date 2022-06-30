// channel(1)
package main

import (
	"fmt"
	_ "fmt"
	"time"
	_ "time"
)

func work1(v chan int) {
	fmt.Println("Work1 : S --> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work1 : E --> ", time.Now())
	v <- 1 // 송신
}

func work2(v chan int) {
	fmt.Println("Work2 : S --> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work2 : E --> ", time.Now())
	v <- 2 // 송신
}

func main() {
	// Channel
	// Goroutine 간의 상호 정보(데이터) 교환 및 실행 흐름 동기화 위해 사용 : Channel(동기식, 참조형)
	// 실행 흐름 제어 가능 (동기, 비동기) -> 일반 변수로 선언 후 사용 가능
	// 데이터 전달 자료형 선언 후 사용(지정된 타입 만 주고 받을 수 있음)
	// interface{} 전달 하면 모든 데이터 타입 전송 및 수신 가능
	// 값을 전달 (복사 후 : bool, int 등), Pointer(슬라이스, 맵) 등을 전달 시에는 주의 ! > 동기화 사용 (Mutex)
	// 멀티프로세싱 처리에서 교착상태(Deadlock, 경쟁) 주의 !
	// <-, -> ( 채널 <- 데이터 : 송신) , ( 채널 -> 데이터 : 수신
	// Channel 은 기본적으로 동기

	// 예제 1.
	fmt.Println("Main : S --->", time.Now())

	// int 형 Channel 생성
	// 방법 1. 정석적인 방법
	// var c chan int
	// c = make(chan int)

	// 방법 2. 짧은 선언
	v := make(chan int)

	go work1(v)
	go work2(v)

	<-v
	<-v
	// time.Sleep()

	fmt.Println("Main : E --->", time.Now())
}
