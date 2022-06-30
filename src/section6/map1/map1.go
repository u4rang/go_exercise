// 자료형 : Map(1)
package main

import "fmt"

func main() {
	// Map
	// - 해시테이블, 딕셔너리(Python)
	// - (Key, Value) 로 자료 저장
	// - 래퍼런스 타입 (참조 값 전달)
	// - 참조 타입이므로 비교 연산자 사용 불가
	// - 특징
	// -  1. 참조타입(Key) 로 사용 불가능, 값(Value)로 모든 타입 사용가능
	// - make 함수 및 축약(리터럴)로 초기화 가능
	// - 순서 없음

	// 예제 1. 선언
	var map1 map[string]int = make(map[string]int) // 문법 상 정석
	var map2 = make(map[string]int)                // 자료형 생략
	map3 := make(map[string]int)                   // 리터럴형, 축약형

	fmt.Printf("map1 : %v\n", map1)
	fmt.Printf("map2 : %v\n", map2)
	fmt.Printf("map3 : %v\n", map3)
	fmt.Println()

	// 예제 2.
	map4 := map[string]int{} // Json 형태
	map4["apple"] = 25
	map4["banana"] = 40
	map4["oragne"] = 33

	map5 := map[string]int{
		"apple":  15,
		"banana": 20,
		"oragne": 30,
	}

	fmt.Printf("map4 : %v\n", map4)
	fmt.Printf("map5 : %v\n", map5)
	fmt.Println()

	map6 := make(map[string]int, 10)
	map6["apple"] = 25
	map6["banana"] = 40
	map6["oragne"] = 33

	fmt.Printf("map6 : %v\n", map6)
	fmt.Printf("map6[\"apple\"] : %d\n", map6["apple"])

}
