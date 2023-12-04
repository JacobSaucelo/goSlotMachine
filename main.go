package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var test = [][]int{
	{0, 1, 2},
	{1, 2, 2},
	{1, 0, 2},
}

func main() {

	slot1 := NewSlotMachine()

	for {
		reader := bufio.NewReader(os.Stdin)

		playerInput, _ := reader.ReadString('\n')
		if strings.ToLower(strings.TrimSpace(playerInput)) == "e" {
			break
		}

		slot1.Spin()
		if slot1.CheckWin() {
			fmt.Println("You Won")
		} else {
			fmt.Println("Unlucky try again?")
		}
		slot1.Display()

	}

}
