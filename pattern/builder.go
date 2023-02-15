package pattern

//! check again
import "fmt"

type Director struct {
	builder IBuilder
}

func (d Director) MakeOperetion(num int) int {
	d.builder.SetNum(num)
	d.builder.Operation()
	return d.builder.GetResult()
}

type IBuilder interface {
	SetNum(int)
	Operation()
	GetResult() int
}

func GetBuilder(num int) IBuilder {
	if num == 2 {
		return &pow2{}
	}
	return &pow3{}
}

type pow2 struct {
	value int
}

func (p *pow2) SetNum(num int) {
	p.value = num
}
func (p *pow2) Operation() {
	p.value = p.value * p.value
}
func (p *pow2) GetResult() int {
	return p.value
}

type pow3 struct {
	value int
}

func (p *pow3) SetNum(num int) {
	p.value = num
}
func (p *pow3) Operation() {
	p.value = p.value * p.value * p.value
}
func (p *pow3) GetResult() int {
	return p.value
}

func main() {
	bPow2 := GetBuilder(2)
	bPow3 := GetBuilder(3)

	director := Director{bPow2}
	fmt.Println("director{pow2}: ", director.MakeOperetion(2))

	director = Director{bPow3}
	fmt.Println("director{pow3}: ", director.MakeOperetion(2))

}
