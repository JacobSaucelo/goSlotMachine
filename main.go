package main

import "fmt"

var test = [][]int{
	{0, 1, 2},
	{1, 2, 2},
	{1, 0, 2},
}

func main() {
	slot1 := NewSlotMachine()

	slot1.Spin()
	fmt.Println(slot1)
}
