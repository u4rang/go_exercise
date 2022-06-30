// 데이터 타입 : Bool(2)

package main

import "fmt"

func main() {
	fmt.Println("true && true :", true && true)
	fmt.Println("true && false :", true && false)
	fmt.Println("false && false :", false && false)

	fmt.Println("true || true :", true || true)
	fmt.Println("true || false :", true || false)
	fmt.Println("false || false :", false || false)
	fmt.Println("!true :", !true)
	fmt.Println("!false :", !false)

	num1 := 15
	num2 := 30

	fmt.Println("num1 < num2 : ", num1 < num2)
	fmt.Println("num1 > num2 : ", num1 > num2)
	fmt.Println("num1 <= num2 : ", num1 <= num2)
	fmt.Println("num1 >= num2 : ", num1 >= num2)
	fmt.Println("num1 == num2 : ", num1 == num2)
	fmt.Println("num1 != num2 : ", num1 != num2)
}
