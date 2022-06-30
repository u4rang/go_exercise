// For문(2)

package main

import "fmt"

func main() {
	// 예제 1
	sum1 := 0

	for i := 0; i <= 100; i++ {
		// sum1 += 1 // sum = sum + 1
		sum1 += i
	}
	fmt.Println("ex1 : ", sum1)

	// 예제 2
	sum2, i := 0, 0

	for i <= 100 {
		sum2 += i
		i++
		// j := i++ // Go에서 후치연산은 반환값 없음 - 에러 발생
	}
	fmt.Println("ex2 : ", sum2)

	// 예제 3, While 형태와 비슷
	sum3, j := 0, 0
	for {
		if j > 100 {
			break
		}
		sum3 += j
		j++
	}
	fmt.Println("ex3 : ", sum3)

	// 예제 4.
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println("ex4 : ", i, j)
	}
}
