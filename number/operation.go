package number

type Operation struct{}

func newOperation() Operation {
	return Operation{}
}

func (o Operation) Add(a, b int) int {
	return a + b
}
