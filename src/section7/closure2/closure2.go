// Closure(2)
package main

import (
	"fmt"
)

func main() {
	// Closure

	// 예제 1.
	cnt := increaseCnt()

	fmt.Println("increaseCnt : ", cnt())
	fmt.Println("increaseCnt : ", cnt())
	fmt.Println("increaseCnt : ", cnt())
	fmt.Println("increaseCnt : ", cnt())
	fmt.Println("increaseCnt : ", cnt())
	fmt.Println("increaseCnt : ", cnt())
	fmt.Println("increaseCnt : ", cnt())
	fmt.Println("increaseCnt : ", cnt())
	fmt.Println("increaseCnt : ", cnt())
	fmt.Println("increaseCnt : ", cnt())

	// cnt = nil

	// cnt = increaseCnt()

	fmt.Println("increaseCnt : ", cnt())
	fmt.Println("increaseCnt : ", cnt())
	fmt.Println("increaseCnt : ", cnt())
	fmt.Println("increaseCnt : ", cnt())
	fmt.Println("increaseCnt : ", cnt())
	fmt.Println("increaseCnt : ", cnt())
	fmt.Println("increaseCnt : ", cnt())
	fmt.Println("increaseCnt : ", cnt())
	fmt.Println("increaseCnt : ", cnt())
	fmt.Println("increaseCnt : ", cnt())
}

func increaseCnt() func() int {
	n := 0 // 지역변수 (캡쳐 됨)
	return func() int {
		n += 1
		return n
	}

}
