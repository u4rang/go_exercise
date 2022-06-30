// 자료형 : Map(4)
package main

import "fmt"

func main() {
	// map 조회 할 경우 주의 할 점

	// 예제 1.
	// 데이터 기본값
	// - int: 0
	// - string: ""
	// - boolean: false
	map1 := map[string]int{
		"apple":  15,
		"banana": 20,
		"oragne": 25,
		"lemon":  0,
	}

	value1, ok1 := map1["lemon"]
	value2 := map1["kiwi"]
	value3, ok2 := map1["kiwi"] // 두번 쨰 리턴 값으로 키 존재 유무 확인

	fmt.Println("value1 : ", value1)
	fmt.Println("value2 : ", value2)
	fmt.Println("value3 : ", value3)
	fmt.Println("ok1 : ", ok1)
	fmt.Println("ok2 : ", ok2)

	// 예제 2.
	if value, ok := map1["kiwi"]; ok {
		fmt.Println("ok is true and kiwi is ", value)
	} else {
		fmt.Println("ok is false and kiwi is not exist!")
	}

	if value, ok := map1["banana"]; ok {
		fmt.Println("ok is true and banana is ", value)
	} else {
		fmt.Println("ok is false and banana is not exist!")
	}

	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("ok is false and kiwi is not exist!")
	}
}
