package example

type Example struct{}

func (e *Example) Greeting() string {
	return "Hello"
}

func (e *Example) Add(a, b int) int {
	return a + b
}

func (e *Example) Sub(a, b int) int {
	return a - b
}

func (e *Example) Mul(a, b int) int {
	return a * b
}

func (e *Example) Div(a, b int) int {
	return a / b
}
