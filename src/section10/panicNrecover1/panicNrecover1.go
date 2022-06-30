// Panic & Recover(1)
package main

import (
	_ "errors"
	"fmt"
	_ "fmt"
	_ "log"
)

func main() {
	// Panic & Recover
	// 컴파일 언어는 컴파일 시 문법적인 오류가 발생하면 에러 처리하여 에러를 원천적으로 처리
	// 동적 타입 언어(golang) 은 런타임 시에 Panic & Recover로 에러 처리
	// 사용자가 에러 발생 가능
	// Panic 함수는 호출 즉시,
	// 해당 메소드를 즉시 중지시키고,
	// defer 함수를 호출하고 자기 자신을 호출한 곳으로 리턴
	// 런타임 이외에 사용자가 코드 흐름에 따라 에러를 발생 시킬 때 중요!
	// 문법적이 에러는 아니지만, 논리적인 코드 흐름에 따른 에러 발생 처리 가능

	fmt.Println("Start Main")
	panic("Error occurred : user stopped !") // 방법 1
	// log.panic("Error occurred : user stopped !") // 방법 2
	fmt.Println("End Main") // 실행 불가
}
