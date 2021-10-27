package example

type Example struct{}

func (e *Example) Greeting() string {
	return "Hello"
}

func (e *Example) Echo(s string) string {
	return s
}
