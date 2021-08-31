package arithmetic

// 두 숫자의 사칙연산 계산 제공 패키지(1)

// X, Y 2개의 Interger 구조체

type Numbers struct {
	X int
	Y int
}

// Plus x,y 합을 계산해서 반환
func (o *Numbers) Plus() int {
	return o.X + o.Y
}

// Minus x,y 차를 계산해서 반환
func (o *Numbers) Minus() int {
	return o.X - o.Y
}

// Multi x,y 곱을 계산해서 반환
func (o *Numbers) Multi() int {
	return o.X * o.Y
}

// Divide x,y 나누기를 계산해서 반환
func (o *Numbers) Divide() int {
	return o.X / o.Y
}
