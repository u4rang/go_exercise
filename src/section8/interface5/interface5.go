// 인터페이스 기본(5)

package main

import (
	"fmt"
	_ "fmt"
)

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

func printValue(s interface{}) {
	fmt.Println("printValue : ", s)
}

func main() {
	// 인터페이스 활용 - 빈 인터페이스
	// 함수 내에서 어떠한 타입이라도 유연하게 매개변수로 받을 수 있다. (만능) -> 모든 타입 지정 가능

	dog := Dog{"poll", 10}
	cat := Cat{"tom", 5}

	printValue(dog)
	printValue(cat)

	printValue(15)
	printValue("Hello 빈 인터페이스")
	printValue(12.345)
	printValue([]Dog{})
	printValue([5]Dog{})

}
