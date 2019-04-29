package main

import (
	"fmt"
	"sync"
)

func main() {
	a := 0
	times := 10000
	c := make(chan bool)

	var m sync.Mutex

	for i := 0; i < times; i++ {
		go func() {
			m.Lock()
			a++
			m.Unlock()
			c <- true
		}()
	}

	for i := 0; i < times; i++ {
		<-c
	}
	fmt.Printf("a = %d\n", a)
}
