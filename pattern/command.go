package pattern

import "fmt"

type Button struct {
	command ICommand
}

func (b Button) PressButton() {
	b.command.execute()
}

type ICommand interface {
	execute()
}

type IDevice interface {
	start()
	end()
}
type StartCommand struct {
	device IDevice
}

func (s StartCommand) execute() {
	s.device.start()
}

type EndCommand struct {
	device IDevice
}

func (s EndCommand) execute() {
	s.device.end()
}

type Counter struct {
	numParam int
	status   bool
}

func (c Counter) start() {
	c.numParam++
	c.status = true
	fmt.Println("Counter status:", c.status)
}
func (c Counter) end() {
	c.numParam++
	c.status = false
	fmt.Println("Counter status:", c.status)
}

func main() {
	counter := &Counter{}
	startCommand := &StartCommand{counter}
	endCommand := &EndCommand{counter}

	startButton := Button{startCommand}
	startButton.PressButton()

	endButton := Button{endCommand}
	endButton.PressButton()
}
