// GO 특징(1)

package main

import "fmt"

func main() {
	// GO
	// - 모호하거나, 애매한 문법 금지
	// -- 후위 연산자는 허용(i++), 전위 연산자는 불가(++i)
	// -- 증감 연산자는 반화값 없음 sum := i++
	// -- 포인터(Pointer) 허용, 하지만 연산은 비허용
	// -- 주석( //, /* ~~~ */)

	// 예제 1.
	sum, i := 0, 0

	for i <= 100 {
		// sum += i++
		sum += i
		i++
		// ++i 예외 발생
	}

	fmt.Println("ex 1 : ", sum)

}
