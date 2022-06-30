// 자료형 : slice(4)
package main

import (
	"fmt"
	"sort"
)

func main() {
	// Slice 추출 및 정렬
	// Slice[i, j] -> i 부터 j-1 까지 추출
	// Slice[i:] -> i 부터 마지막까지 추출
	// Slice[:j] -> 처음부터 j-1 까지 추출
	// Slice[:] -> 처음부터 마지막까지 추출

	// 예제 1. 추출
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("slice1[:]", slice1[:])
	fmt.Println("slice1[0:]", slice1[0:])
	fmt.Println("slice1[:5]", slice1[:5])
	fmt.Println("slice1[0:len(slice1)]", slice1[0:len(slice1)])
	fmt.Println("slice1[:3]", slice1[:3])
	fmt.Println("slice1[3:]", slice1[3:])
	fmt.Println("slice1[1:3]", slice1[1:3])
	fmt.Println("slice1[3:3]", slice1[3:3])

	// 예제 2. 정렬
	// sort 패키지 : https://golang.org/pkg/sort

	slice2 := []int{3, 6, 9, 1, 2, 4, 5, 7, 8}
	slice3 := []string{"b", "c", "d", "a", "c", "e"}

	fmt.Printf("slice2 :%b %v\n", sort.IntsAreSorted(slice2), slice2) // 정렬 확인

	sort.Ints(slice2) // 정렬

	fmt.Printf("slice2 :%b %v\n", sort.IntsAreSorted(slice2), slice2) // 정렬 확인

	fmt.Printf("slice3 :%b %v\n", sort.StringsAreSorted(slice3), slice3) // 정렬 확인

	sort.Strings(slice3) // 정렬

	fmt.Printf("slice3 :%b %v\n", sort.StringsAreSorted(slice3), slice3) // 정렬 확인
}
