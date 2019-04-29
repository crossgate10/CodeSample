package main

import (
	"fmt"
)

func main() {
	ns := []string{"Jade", "Kavin", "Julia"}
	nss := ns[1:]

	A(ns)
	A(nss)

	fmt.Printf("%p\n", ns)
	fmt.Printf("%p\n", nss)

	fmt.Println("---------------------")

	B(ns...)
	B(nss...)

	fmt.Printf("%p\n", ns)
	fmt.Printf("%p\n", nss)
}

func A(args ...interface{}) {
	fmt.Printf("%p\n", args)
}

func B(args ...string) {
	fmt.Printf("%p\n", args)
}
