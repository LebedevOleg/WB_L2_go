package pattern

import "fmt"

type ILigth interface {
	LightStatus()
}
type Lamp struct {
	state ILigth
}

func (l *Lamp) LightStatus() {
	l.state.LightStatus()
}

func (l *Lamp) SwitchLight(state ILigth) {
	l.state = state
	l.LightStatus()
}
func NewLamp() Lamp {
	return Lamp{Lamp_off{}}
}

type Lamp_off struct{}

func (l Lamp_off) LightStatus() {
	fmt.Println("Свет выключен")
}

type Lamp_on struct{}

func (l Lamp_on) LightStatus() {
	fmt.Println("Свет включен")
}

func main() {
	lamp := NewLamp()
	lamp.LightStatus()
	lamp.SwitchLight(Lamp_on{})
	lamp.SwitchLight(Lamp_off{})
}
