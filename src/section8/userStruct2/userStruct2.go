//  사용자 정의 타입(2)

package main

import (
	"fmt"
	_ "fmt"
)

type cnt int

func main() {
	// 기본 자료형 사용자 정의 타입 활용 가능

	// 예제 1.
	a := cnt(5)

	fmt.Println("a : ", a)

	// 예제 2. 추천방법
	var b cnt = 5

	fmt.Println("b : ", b)

	// 예제 3.
	// testConvertI(a)	// 에러 발생 중요 ! : 사용자 정의 타입 <-> 기본 타입 : 매개 변수 전달 시 변환해야 사용 가능(int(변수))
	testConvertT(int(a)) // 명시적 형변환 시 가능
	testConvertD(a)
}

func testConvertT(i int) {
	fmt.Println("testConvertT : (Default Type) : ", i)
}

func testConvertD(i cnt) {
	fmt.Println("testConvertD : (User Defined Type) : ", i)
}
