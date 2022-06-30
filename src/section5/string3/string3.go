// 데이터 타입 : 문자열 연산(1)
package main

import "fmt"

func main() {
	// 문자열 연산
	// 추출, 비교, 조합(결합)

	// 예제 1. 추출
	var str1 string = "Golang"
	var str2 string = "World"

	fmt.Println("str1[0:2], str1[0] : ", str1[0:2], str1[0])
	fmt.Println("str2[3:], str2[0] : ", str2[3:], str2[0])
	fmt.Println("str2[:4]", str2[:4])
	fmt.Println("str2[1:3]", str2[1:3])
}
