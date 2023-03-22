package pattern

import "fmt"

type Handler interface {
	SendRequest(message int) string
}

type CHandlerA struct {
	next Handler
}

func (a *CHandlerA) SendRequest(m int) string {
	if m == 1 {
		return "Обработчик А"
	}
	if a.next != nil {
		return a.next.SendRequest(m)
	}
	return "Ничего"

}

type CHandlerB struct {
	next Handler
}

func (b *CHandlerB) SendRequest(m int) string {
	if m == 2 {
		return "Обаработчик B"
	}
	if b.next != nil {
		return b.next.SendRequest(m)
	}
	return "Ничего"
}

type CHandlerС struct {
	next Handler
}

func (c *CHandlerС) SendRequest(m int) string {
	if m == 2 {
		return "Обаработчик B"
	}
	if c.next != nil {
		return c.next.SendRequest(m)
	}
	return "Ничего"
}

func main() {
	handler := &CHandlerA{&CHandlerB{&CHandlerС{}}}

	fmt.Println(handler.SendRequest(2))
}
