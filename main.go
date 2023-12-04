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
	fmt.Println("Press enter to start")

	for {
		reader := bufio.NewReader(os.Stdin)

		playerInput, _ := reader.ReadString('\n')
		if strings.ToLower(strings.TrimSpace(playerInput)) == "e" {
			break
		}
		if strings.ToLower(strings.TrimSpace(playerInput)) == "r" {
			slot1.Credits += 100
		}

		slot1.Spin()
		status, multipier := slot1.CheckWin()
		if status {
			if multipier == 6 {
				slot1.Credits *= 3
				slot1.Display("DIAMOND: CREDITS MULTIPLIED")
			} else {
				slot1.Display("You Won")
			}
		} else {
			slot1.Display("Unlucky try again?")
		}
	}

}
