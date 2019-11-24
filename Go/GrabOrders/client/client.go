package main

import (
	"GrabOrders/pool/repo"
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	routines := make([]int, 10)
	for i := range routines {
		routines[i] = i+1
		go func(csID int64) {
			for {
				orderID := repo.GetFirstUngrabOrders(csID)
				err := repo.GrabOrder(csID, orderID)
				if err == nil {
					c <- fmt.Sprintf("Customer %d grab order %d \n", csID, orderID)
					time.Sleep(1 * time.Second)
				}
			}
		}(int64(routines[i]))
	}

	for {
		select {
		case s := <-c:
			fmt.Printf(s)
		}
	}
}
