package number

type Num struct{}

func (n *Num) Add(a, b int) int {
	return a + b
}

func (n *Num) AddSubstract(a, b int) (int, int) {
	return a + b, a - b
}
