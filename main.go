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

var chips uint32 = 10

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

		slot1.Spin(chips)
		status, multipier := slot1.CheckWin()
		if status {
			if multipier == 6 {
				multipliedReward := slot1.Credits * 3
				if multipliedReward > 1000 {
					slot1.Credits += 1000
					slot1.Display("DIAMOND: +1000 REACHED MAX WIN!")
				} else {
					slot1.Display("DIAMOND: CREDITS MULTIPLIED!")
					slot1.Credits *= 3
				}
			} else {
				slot1.Display("You Won")
			}
		} else {
			slot1.Display("Unlucky try again?")
		}
	}

}
