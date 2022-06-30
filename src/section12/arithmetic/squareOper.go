// Package -  두 숫자의 연산 계산 제공 패키지(2)
package arithmetic

// x, y 제곱의 합을 리턴
func (num *Numbers) SquarePlus() int {
	return (num.X * num.X) + (num.Y * num.Y)
}

// x, y 제곱의 차를 리턴
func (num *Numbers) SquareMinus() int {
	return (num.X * num.X) - (num.Y * num.Y)
}
