// 라이브러리 접근제어(1)

package lib

import "fmt"

var num int32
var Num2 int32

// 변수 초기화
func init() {
	num = 9
	Num2 = 100
	fmt.Println("lib Package > init start!")
	fmt.Println("lib Packe > num : ", num)
	fmt.Println("lib Packe > Num2 : ", Num2)
}

func CheckNum(c int32) bool {
	return c > 10
}
