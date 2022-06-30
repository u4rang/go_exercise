// 인터페이스 기본(1)

package main

import (
	"fmt"
	_ "fmt"
)

type Dog struct {
	name   string
	weight int
}

// bite 메소드 구형
func (d Dog) bite() {
	fmt.Println(d.name, "bite!")
}

//  동물의 행동 인터페이스 선언
type Behavior interface {
	bite()
}

func main() {
	// 인터페이스

	// 예제 1.
	dog1 := Dog{"poll", 1}

	var inter1 Behavior
	inter1 = dog1
	inter1.bite()

	// dog1.bite()

	// 예제 2.
	dog2 := Dog{"merry", 12}
	inter2 := Behavior(dog2)
	inter2.bite()

	// 예제 3.
	inters := []Behavior{dog1, dog2}

	// 인덱스 형태로 설정
	for i, _ := range inters {
		inters[i].bite()
	}

	// 값 형태로 실행(인터페이스)
	for _, val := range inters {
		val.bite()
	}
}
