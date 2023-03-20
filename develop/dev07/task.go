package main

import (
	"fmt"
	"time"
)

func Single(channels ...<-chan interface{}) <-chan interface{} {
	for {
		for _, v := range channels {
			select {
			case <-v:
				return v
			default:
				continue
			}
		}
	}
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-Single(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(20*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
		sig(10*time.Second),
	)

	fmt.Printf("fone after %v", time.Since(start))

}
