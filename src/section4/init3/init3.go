// go 초기화 (#)

package main

import (
	"fmt"
	"section4/lib"
)

var num int32

// 변수 초기화
func init() {
	num = 30
	fmt.Println("Main Package > init start!")
}

func main() {
	fmt.Println("10 보다 큰 수 ? : ", lib.CheckNum(num))
	fmt.Println("Main Packe > num : ", num)
	fmt.Println("Main Packe > Num2 : ", lib.Num2)
}
