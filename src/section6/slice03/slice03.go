// 자료형 : slice(3)
package main

import (
	"fmt"
)

func main() {
	// Slice 추가 및 병합

	// 예제 1.
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8, 9, 10}
	s3 := []int{11, 12, 13, 14, 15}

	s1 = append(s1, 6, 7)
	s2 = append(s1, s2...)      // Slice 를 사용할 경우 사용
	s3 = append(s3, s2[0:3]...) // 추출 후 병합

	fmt.Printf("s1 : %v\n", s1)
	fmt.Printf("s2 : %v\n", s2)
	fmt.Printf("s3 : %v\n", s3)

	// 예제 2.
	s4 := make([]int, 0, 5)
	for i := 0; i < 15; i++ {
		s4 = append(s4, i)
		fmt.Printf("s4 len : %d , cap : %d , data : %v\n", len(s4), cap(s4), s4)
	}

}
