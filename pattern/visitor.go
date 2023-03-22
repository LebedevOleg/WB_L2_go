package pattern

import "fmt"

type Visitor interface {
	VisitPlace1(p *Place1) string
	VisitPlace2(p *Place2) string
}

type Place interface {
	Accept(v Visitor) string
}

type People struct{}

func (human *People) VisitPlace1(p *Place1) string {
	return p.VisitPlace()
}
func (human *People) VisitPlace2(p *Place2) string {
	return p.VisitPlace()
}

type Place1 struct{}

func (p *Place1) Accept(v Visitor) string {
	return v.VisitPlace1(p)
}
func (p *Place1) VisitPlace() string {
	return "Был в месте 1"
}

type Place2 struct{}

func (p *Place2) Accept(v Visitor) string {
	return v.VisitPlace2(p)
}
func (p *Place2) VisitPlace() string {
	return "Был в месте 2"
}

func main() {
	city := make([]Place, 0)
	city = append(city, &Place1{})
	city = append(city, &Place2{})
	for _, v := range city {
		fmt.Println(v.Accept(&People{}))
	}
}
