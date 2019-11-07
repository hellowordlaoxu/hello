package calc

func Div(a, b int) int {
	if b == 0 {
		panic("除数不能为零")
	}
	return a / b
}
