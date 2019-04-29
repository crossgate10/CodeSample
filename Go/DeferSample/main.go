package main

import "fmt"

func main() {
	second()
}

func calc(index, a, b int) int{
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func second(){
	a := 1
	b := 2
	defer calc(10, a, calc(20, a, calc(30, a, b)))
	a = 0
	defer calc(40, a, calc(50, a, calc(60, a, b)))
	b = 5
}

//--------------------------------------------------

func A(){
	fmt.Println("In A")
}

func B(){
	fmt.Println("In B")
}

func C(){
	fmt.Println("In C")
}

func first(){
	defer A()
	defer B()
	defer C()
	panic("something error.")
}

