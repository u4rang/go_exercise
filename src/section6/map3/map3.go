// 자료형 : Map(3)
package main

import "fmt"

func main() {
	// map 값 변경 및 수정

	// 예제 1.
	map1 := map[string]string{
		"daum":   "https://www.daum.net",
		"naver":  "https://www.naver.com",
		"google": "https://www.google.com",
		"home1":  "https://www.home1.com",
	}

	fmt.Println("map1 : ", map1)

	map1["home2"] = "https://www.home2.com" // 추가

	fmt.Println("map1 : ", map1)

	map1["home2"] = "https://www.home2_modified.com" // 수정

	fmt.Println("map1 : ", map1)

	// 예제 2. 삭제
	delete(map1, "home2")

	fmt.Println("map1 : ", map1)

	delete(map1, "google")

	fmt.Println("map1 : ", map1)
}
