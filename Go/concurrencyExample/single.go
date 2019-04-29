package main

import (
	"fmt"
)

func singleRun() {
	var theMine, foundOre, minedOre []string
	theMine = []string{"rock", "ore", "ore", "rock", "ore"}
	foundOre = finder(theMine)
	minedOre = miner(foundOre)
	smelter(minedOre)
}

func finder(theMine []string) []string {
	var foundOre []string
	for i := range theMine {
		if theMine[i] == "ore" {
			foundOre = append(foundOre, "foundOre")
		}
	}
	return foundOre
}

func miner(foundOre []string) []string {
	var minedOre []string
	for range foundOre {
		minedOre = append(minedOre, "minedOre")
	}
	return minedOre
}

func smelter(minedOre []string) {
	var smeltedOre []string
	for range minedOre {
		smeltedOre = append(smeltedOre, "smeltedOre")
	}
	fmt.Printf("%v", smeltedOre)
}
