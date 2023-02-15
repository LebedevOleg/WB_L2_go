package pattern

import "fmt"

type SystemA struct {
	valueA  string
	activeA bool
}

func (a SystemA) StartSystemA(value string) string {
	a.valueA = value
	a.activeA = true
	return a.valueA
}

type SystemB struct {
	valueB  int
	activeB bool
}

func (b SystemB) StartSystemB(value int) int {
	b.valueB = value
	b.activeB = true
	return b.valueB
}

type FacadeAB struct {
	a           SystemA
	b           SystemB
	facadeValue string
}

func (f FacadeAB) StartFacade(text string, num int) string {
	f.facadeValue = f.a.StartSystemA(text)
	f.facadeValue = fmt.Sprint(f.facadeValue, " ", f.b.StartSystemB(num))
	return f.facadeValue
}

func main() {
	facade := FacadeAB{}
	fmt.Println(facade.StartFacade("я веселый текст", 666))
}
