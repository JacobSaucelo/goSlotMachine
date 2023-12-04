package main

import (
	"fmt"
	"math/rand"
)

var test = [][]int{
	{0, 1, 2},
	{1, 2, 2},
	{1, 0, 2},
}

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(getRandomNumber())
	}

}

func getRandomNumber() int {

	return rand.Intn(10-1) + 1
}
