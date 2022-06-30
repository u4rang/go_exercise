// Function(2)
package main

import (
	"fmt"
)

func sum(i int, f func(int, int)) {
	f(i, 10)
}

func add(a, b int) {
	fmt.Println("add : ", a+b)
}

func multiValue(i int) {
	i = i * 10
}

func multiReference(i *int) {
	*i = *i * *i
}

func main() {
	// Function
	// Function(Callback) 전달, 참조 전달(call by reference), 값 전달(call by value)

	// 예제 1. 함수 전달
	sum(100, add)

	// 예제 2. call by value
	a := 100
	multiValue(a)
	fmt.Println("multiValue", a)

	// 예제 2. call by reference
	b := 100
	multiReference(&b)
	fmt.Println("multiReference", b)

}
