// 인터페이스 기본(4)

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

// 구조체 Dog 메소드 구현
func (d Dog) run() {
	fmt.Println(d.name, " - Dog run")
}

// 구조체 Cat 메소드 구현
func (c Cat) run() {
	fmt.Println(c.name, " - Cat run")
}

func act(animal interface{ run() }) { // 익명 인터페이스(타입 정의  x)
	animal.run()
}

func main() {
	// 인터페이스
	// 익명 인터페이스 사용 예제 (즉시 사용 후 사용)

	// 예제 1.
	dog := Dog{"poll", 10}
	cat := Cat{"tom", 5}

	act(dog)
	act(cat)
}
