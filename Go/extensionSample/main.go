package main

import "fmt"

type CustomString string

func (c CustomString) PrintAndAddSufix(s string) string {
	fmt.Println(s)
	return "'" + s + "' already print!!!"
}

func main() {
	var s CustomString
	ss := s.PrintAndAddSufix("TEST")
	fmt.Println(ss)
}
