package main

import (
	"fmt"
	"sort"
)

type Sequence []int

// Method for printing - sorts the elements before printing
func (s Sequence) String() string {
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}

func main(){
	serialNumbers := Sequence([]int{9, 6, 8, 10, 4, 7, 3, 2, 5, 1, 0})

	fmt.Println(serialNumbers.String())
}
