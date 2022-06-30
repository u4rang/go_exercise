// 인터페이스 고급(2)

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 인터페이스 고급 - Type Assertion
	// 타입 변환
	// 실행(런타임) 시에는 인터페이스에 할당된 변수는 실제 타입으로 변환 후 사용해야하는 경우
	// interfaceVal.(type)

	// 예제 1.
	var a interface{} = 15
	b := a
	c := a.(int)
	// d := a.(float64)	// 에러 발생 - 최초의 데이터 타입으로만 변환 가능

	fmt.Println("a : ", a)
	fmt.Println("a type of : ", reflect.TypeOf(a))
	fmt.Println("b : ", b)
	fmt.Println("b type of : ", reflect.TypeOf(b))
	fmt.Println("c : ", c)
	fmt.Println("c type of : ", reflect.TypeOf(c))
	//fmt.Println("d : ", d)	// 에러발생
	fmt.Println()

	// 예제 2. 저정된 실제 타입 검사
	if v, ok := a.(int); ok {
		fmt.Println("value : ", v, " - isOk :", ok)

	}

}
