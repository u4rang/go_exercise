// Panic & Recover(2)
package main

import (
	_ "errors"
	"fmt"
	_ "fmt"
	_ "log"
)

func runFunc() {
	fmt.Println("runFunc Function Start")

	// Recover 에러 처리
	defer func() {
		s := recover()
		fmt.Println("Error Message : ", s)
	}()

	panic("Error occurred!")

	fmt.Println("runFunc Function End") // 호출 안됨

}

func main() {
	// Panic & Recover
	// golang recover() function
	// - Panic으로 발생한 에러를 복구 후 코드 재실행(프로그램 종료 되지 않는다)
	// - 즉, 코드 흐름을 정상상태로 복구하는 기능
	// - Panic에서 설정한 메시지를 받아 올 수 있다.

	runFunc()

	fmt.Println("Hello golang!")
}
