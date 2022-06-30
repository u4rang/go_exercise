// 데이터 타입 : 문자열(2)
package main

import "fmt"

func main() {
	// 문자열 표현
	// 문자열 : UTF-8 인코딩 (유니코드 문자 집합) -> 바이트 수 주의!

	// 예제 1.
	var str1 string = "Golang"
	var str2 string = "World"
	var str3 string = "고프로그래밍"

	fmt.Println("str1 :", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Println("str2 :", str2[0], str2[1], str2[2], str2[3], str2[4])
	fmt.Println("str3 :", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5])

	// 예제 2.
	fmt.Printf("str1 : %c %c %c %c %c %c \n", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Printf("str2 : %c %c %c %c %c \n", str2[0], str2[1], str2[2], str2[3], str2[4])
	fmt.Printf("str3 : %c %c %c %c %c %c \n", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5]) // 한글 꺠짐

	conStr := []rune(str3)                                                                                        // 한글은 배열형 rune 으로 접근
	fmt.Printf("conStr : %c %c %c %c %c %c \n", conStr[0], conStr[1], conStr[2], conStr[3], conStr[4], conStr[5]) // 한글 정상 출력

	// 예제 3.
	for i, char := range str1 {
		fmt.Printf("%c(%d)\t", char, i)
	}
	fmt.Println()

	for i, char := range str2 {
		fmt.Printf("%c(%d)\t", char, i)
	}
	fmt.Println()

	for i, char := range str3 {
		fmt.Printf("%c(%d)\t", char, i)
	}
	fmt.Println()

	for i, char := range conStr {
		fmt.Printf("%c(%d)\t", char, i)
	}
	fmt.Println()

}
