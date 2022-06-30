// 자료형 : Map(2)
package main

import "fmt"

func main() {
	// map 조회 및 순회(Iterator)

	// 예제 1.
	map1 := map[string]string{
		"daum":   "https://www.daum.net",
		"naver":  "https://www.naver.com",
		"google": "https://www.google.com",
	}

	fmt.Println("map1[\"daum\"]", map1["daum"])
	fmt.Println("map1[\"google\"]", map1["google"])

	// 예제 2. 순서 없이 랜덤 추출
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	for _, v := range map1 {
		fmt.Println(v)
	}

}
