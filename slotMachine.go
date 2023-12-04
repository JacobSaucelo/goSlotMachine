package main

import (
	"fmt"
	"math/rand"
)

type SymbolType struct {
	id     byte
	icon   string
	reward byte
}

type SlotMachineType struct {
	Reels   [][]SymbolType
	Credits uint32
}

var reward = 0

var Symbols = []SymbolType{
	{
		id:     0,
		icon:   "🍎",
		reward: 10,
	},
	{
		id:     1,
		icon:   "🥭",
		reward: 10,
	},
	{
		id:     2,
		icon:   "🍓",
		reward: 15,
	},
	{
		id:     3,
		icon:   "🍉",
		reward: 10,
	},
	{
		id:     4,
		icon:   "🍇",
		reward: 20,
	},
	{
		id:     5,
		icon:   "🍒",
		reward: 20,
	},
	{
		id:     6,
		icon:   "💎",
		reward: 10,
	},
	{
		id:     7,
		icon:   "🍏",
		reward: 100,
	},
}

func NewSlotMachine() *SlotMachineType {
	return &SlotMachineType{
		Reels:   [][]SymbolType{},
		Credits: 100,
	}
}

func (s *SlotMachineType) Spin() {
	if s.Credits < 10 {
		fmt.Println("You dont have enough credits to continue")
		return
	}

	s.Credits -= 10
	s.Reels = make([][]SymbolType, 3)
	for i := range s.Reels {
		s.Reels[i] = make([]SymbolType, 3)
		for j := range s.Reels[i] {
			s.Reels[i][j] = generateRandomSymbol(0, 7)
		}
	}
}

func (s *SlotMachineType) Display(status string) {
	clearScreen()
	for i := range s.Reels {
		fmt.Print("\n          ")
		fmt.Printf("%s | %s | %s ", s.Reels[i][0].icon, s.Reels[i][1].icon, s.Reels[i][2].icon)
		fmt.Print("\n          ")
		// for j := range s.Reels[i] {
		// 	fmt.Print(strings.TrimSpace(s.Reels[i][j].icon), " | ")
		// }
	}
	fmt.Print("\n")

	fmt.Println("--------------------------------------")
	fmt.Println("|🍎=10|🥭=10|🍉=10|🍓=15|🍇=20|🍒=20|")
	fmt.Println("|💎=(credits x 3)|🍏=100|")
	fmt.Printf("Credits: %d \n\n", s.Credits)
	fmt.Printf("Reward: %d \n\n", reward)
	fmt.Println(status)
	fmt.Println("Press (Enter) key to play more ")
	fmt.Println("Press e to exit")
	fmt.Println("--------------------------------------")
}

func (s *SlotMachineType) CheckWin() (bool, byte) {
	for i := 0; i < 3; i++ {
		if s.Reels[i][0] == s.Reels[i][1] && s.Reels[i][1] == s.Reels[i][2] {
			s.Credits += uint32(s.Reels[i][0].reward)
			reward = int(s.Reels[i][0].reward)
			fmt.Println("you won ", reward)
			return true, s.Reels[i][0].id
		}
		if s.Reels[0][i] == s.Reels[1][i] && s.Reels[1][i] == s.Reels[2][i] {
			s.Credits += uint32(s.Reels[i][0].reward)
			reward = int(s.Reels[0][i].reward)
			fmt.Println("you won ", reward)
			return true, s.Reels[0][i].id
		}
	}

	//? DIAGONAL
	if s.Reels[0][0] == s.Reels[1][1] && s.Reels[1][1] == s.Reels[2][2] {
		s.Credits += uint32(s.Reels[1][1].reward)
		reward = int(s.Reels[1][1].reward)
		fmt.Println("you won ", reward)
		return true, s.Reels[1][1].id
	}
	if s.Reels[0][2] == s.Reels[1][1] && s.Reels[1][1] == s.Reels[2][0] {
		s.Credits += uint32(s.Reels[1][1].reward)
		reward = int(s.Reels[1][1].reward)
		fmt.Println("you won ", reward)
		return true, s.Reels[1][1].id
	}

	reward = 0
	return false, 0
}

func generateRandomSymbol(min, max int) SymbolType {
	//? 	THIS IS BASED ON TIME - GO TOO FAST IT GIVES SAME NUMBER ID
	// source := rand.NewSource(time.Now().UnixNano())
	// rng := rand.New(source)

	return Symbols[rand.Intn(max-min)+1]
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
