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
var winnings = 0

var Symbols = []SymbolType{
	{
		id:     0,
		icon:   "🍎",
		reward: 50,
	},
	{
		id:     1,
		icon:   "🥭",
		reward: 80,
	},
	{
		id:     2,
		icon:   "🍓",
		reward: 95,
	},
	{
		id:     3,
		icon:   "🍉",
		reward: 100,
	},
	{
		id:     4,
		icon:   "🍇",
		reward: 150,
	},
	{
		id:     5,
		icon:   "🍒",
		reward: 150,
	},
	{
		id:     6,
		icon:   "💎",
		reward: 100,
	},
	{
		id:     7,
		icon:   "🍏",
		reward: 250,
	},
}

func NewSlotMachine() *SlotMachineType {
	return &SlotMachineType{
		Reels:   [][]SymbolType{},
		Credits: 100,
	}
}

func (s *SlotMachineType) Spin(chips uint32) {
	if s.Credits < 10 {
		fmt.Println("You dont have enough credits to continue")
		return
	}

	if s.Credits < 100 {
		chips = 10
	}
	if 500 > s.Credits && s.Credits > 100 {
		chips = 40
	}
	if 1000 > s.Credits && s.Credits > 500 {
		chips = 80
	}
	if s.Credits > 1000 {
		chips = 100
	}

	s.Credits -= chips
	s.Reels = make([][]SymbolType, 3)
	for i := range s.Reels {
		s.Reels[i] = make([]SymbolType, 3)
		for j := range s.Reels[i] {
			s.Reels[i][j] = generateRandomSymbol(0, 7)
		}
	}
}

func (s *SlotMachineType) Display(status string) {
	displayCred := 10

	if s.Credits < 100 {
		displayCred = 10
	}
	if 500 > s.Credits && s.Credits > 100 {
		displayCred = 40
	}
	if 1000 > s.Credits && s.Credits > 500 {
		displayCred = 50
	}
	if s.Credits > 1000 {
		displayCred = 100
	}

	clearScreen()
	for i := range s.Reels {
		fmt.Print("\n                 ")
		fmt.Printf("%s | %s | %s ", s.Reels[i][0].icon, s.Reels[i][1].icon, s.Reels[i][2].icon)
		fmt.Print("\n                 ")
		// for j := range s.Reels[i] {
		// 	fmt.Print(strings.TrimSpace(s.Reels[i][j].icon), " | ")
		// }
	}
	fmt.Print("\n")

	fmt.Println("-------------------------------------------------")
	fmt.Println("|🍎 = 50|🥭 = 80|🍓 = 95|🍉= 100|🍇 = 150|🍒 = 150|")
	fmt.Println("|💎 = (credits x 3)|🍏 = 250|")
	fmt.Printf("|spin = %d|\n", displayCred)

	if s.Credits < 100 {
		fmt.Println("|Rock Bottom 🗿💸📉|")
	}
	if 500 > s.Credits && s.Credits > 100 {
		fmt.Println("|Lucky? 💰🍀☘️|")
	}
	if 1000 > s.Credits && s.Credits > 500 {
		fmt.Println("|Thin ICE 🥶🧊⛄|")
	}
	if s.Credits > 1000 {
		fmt.Println("|High stakes 🧈🧈🧈|")
	}
	fmt.Printf("\nCredits: %d \n\n", s.Credits)
	fmt.Printf("Reward: %d \n\n", reward)
	fmt.Println(status)
	fmt.Println("Press (Enter) key to play more ")
	fmt.Println("Press e to exit")
	fmt.Println("-------------------------------------------------")
}

func (s *SlotMachineType) CheckWin() (bool, byte) {
	for i := 0; i < 3; i++ {
		if s.Reels[i][0] == s.Reels[i][1] && s.Reels[i][1] == s.Reels[i][2] {
			s.Credits += uint32(s.Reels[i][0].reward)
			reward = int(s.Reels[i][0].reward)
			return true, s.Reels[i][0].id
		}
		if s.Reels[0][i] == s.Reels[1][i] && s.Reels[1][i] == s.Reels[2][i] {
			s.Credits += uint32(s.Reels[0][i].reward)
			reward = int(s.Reels[0][i].reward)
			return true, s.Reels[0][i].id
		}
	}

	//? DIAGONAL
	if s.Reels[0][0] == s.Reels[1][1] && s.Reels[1][1] == s.Reels[2][2] {
		s.Credits += uint32(s.Reels[1][1].reward)
		reward = int(s.Reels[1][1].reward)
		return true, s.Reels[1][1].id
	}
	if s.Reels[0][2] == s.Reels[1][1] && s.Reels[1][1] == s.Reels[2][0] {
		s.Credits += uint32(s.Reels[1][1].reward)
		reward = int(s.Reels[1][1].reward)
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
