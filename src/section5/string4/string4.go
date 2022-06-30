// 데이터 타입 : 문자열 연산(2)
package main

import "fmt"

func main() {
	// 문자열 연산
	// 추출, 비교, 조합(결합)

	// 예제 2. 비교
	var str1 string = "Golang"
	var str2 string = "World"
	var str3 string = "Golang"

	fmt.Println("str1 == str2 : ", str1 == str2)
	fmt.Println("str1 != str2 : ", str1 != str2)

	// Golang 에서는 문자열 -> ASCII 코드에 대한 사전식 비교
	fmt.Println("str1 > str2 : ", str1 > str2)
	fmt.Println("str1 < str2 : ", str1 < str2)

	fmt.Println("str1 == str3 : ", str1 == str3)

}
