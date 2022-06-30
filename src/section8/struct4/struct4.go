// 구조체(4)

package main

import (
	"fmt"
)

func main() {
	// 구조체 익명 선언 및 활용
	// 예제 1. type 구조체명 타입
	car1 := struct{ name, color string }{"530d", "black"}

	fmt.Println("car1 : ", car1)
	fmt.Printf("car1 : %#v\n", car1)

	// 예제 2. for statements
	cars := []struct{ name, color string }{{"100d", "white"}, {"200d", "blue"}, {"300d", "green"}}

	for _, car := range cars {
		fmt.Printf("(%s, %s)\n", car.name, car.color)
	}
}
