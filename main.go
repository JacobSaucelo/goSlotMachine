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
	slot1.Display()
	fmt.Println("didWin: ", slot1.CheckWin())
	// fmt.Println(slot1)
}
