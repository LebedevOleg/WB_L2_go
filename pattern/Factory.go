package main

import "fmt"

type mathReq interface {
	Operate(int, int) int
}

type Pow struct{}

func (p *Pow) Operate(a int, b int) int {
	res := a
	for b != 1 {
		res = res * a
		b--
	}
	return res
}

type Sum struct{}

func (s Sum) Operate(a, b int) int {
	return a + b
}

func NewMathReqest(reqType string) mathReq {
	switch reqType {
	case "Pow":
		return &Pow{}
	case "Sum":
		return &Sum{}
	default:
		return &Sum{}
	}
}

func main() {
	math := NewMathReqest("Pow")
	fmt.Println(math.Operate(2, 10))
	math = NewMathReqest("Sum")
	fmt.Println(math.Operate(2, 2))

}
