package main

import (
	"fmt"
	"time"
)

func multiRun() {
	var theMine []string
	theMine = []string{"rock", "ore", "ore", "rock", "ore"}
	oreChannel := make(chan string)
	minedOreChan := make(chan string)

	//Finder
	go func(mine []string) {
		for _, item := range mine {
			if item == "ore" {
				oreChannel <- item // send item on oreChannel
			}
		}
	}(theMine)

	// Ore Breaker
	go func() {
		for i := 0; i < 3; i++ {
			foundOre := <-oreChannel // read from oreChannel
			fmt.Println("From Finder: ", foundOre)
			minedOreChan <- "minedOre" // send to minedOreChan
		}
	}()

	// Smelter
	go func() {
		for i := 0; i < 3; i++ {
			minedOre := <-minedOreChan // read from minedOreChan
			fmt.Println("From Miner: ", minedOre)
			fmt.Println("From Smelter: Ore is smelted")
		}
	}()

	<-time.After(time.Second * 5)
}
