// IF문(1)
package main

import "fmt"

func main() {
	// 제어문(IF문)
	// IF 문 : 반드시 Boolean 검사 -> 1, 0 (사용 불가 : 자동 형 변환 불가)
	// 소괄호 사용 불가

	var a int = 20
	b := 20

	// 예제 1
	if a >= 15 {
		fmt.Println("It is upper than 15")
	}

	// 에제 2
	if b >= 25 {
		fmt.Println("It is upper than 25")
	}

	// 에러 발생1
	/*
		if b >= 25
		{
			fmt.Println("Error occur...")
		}
	*/

	// 에러 발생2
	/*
		if b >= 25
			fmt.Println("Error occur...")
	*/

	// 에러 발생3
	/*
		if c := 1; c {
			fmt.Println("Error occur...")
		}
	*/

	// 에러 발생4
	/*
		if c := 40; c >= 35 {
			fmt.Println("Error occur...")
		}

		// Out scope of variable
		fmt.Println(c)
	*/
}
