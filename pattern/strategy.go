package pattern

import "fmt"

type Operator interface {
	Apply(int, int) int
}
type Operation struct {
	operator Operator
}

func (o *Operation) Operate(a int, b int) int {
	return o.operator.Apply(a, b)
}

type Summ struct{}

func (s Summ) Apply(a int, b int) int {
	return a + b
}

type Minus struct{}

func (m Minus) Apply(a int, b int) int {
	return a - b
}

func main() {
	sum := Operation{Summ{}}
	fmt.Println(sum.Operate(2, 2))

	min := Operation{&Minus{}}
	fmt.Println(min.Operate(2, 2))

}
