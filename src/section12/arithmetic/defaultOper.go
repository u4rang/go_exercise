//  Package -  두 숫자의 사칙연산 계산 제공 패키지(1)
package arithmetic

// x, y 2 개의 Integer 구조체
type Numbers struct {
	X int
	Y int
}

// x, y 합을 계산해서 반환
func (num *Numbers) Plus() int {
	return num.X + num.Y
}

// x, y 차를 계산해서 반환
func (num *Numbers) Minus() int {
	return num.X - num.Y
}

// x, y 곱을 계산해서 반환
func (num *Numbers) Multi() int {
	return num.X * num.Y
}

// x, y 나누기를 계산해서 반환
func (num *Numbers) Divide() int {
	return num.X / num.Y
}
