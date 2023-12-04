package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type SymbolType struct {
	id   byte
	icon string
}

type SlotMachineType struct {
	Reels   [][]SymbolType
	Credits uint32
}

var Symbols = []SymbolType{
	{
		id:   0,
		icon: "🍎",
	},
	{
		id:   1,
		icon: "🥭",
	},
	{
		id:   2,
		icon: "🍓",
	},
	{
		id:   3,
		icon: "🍉",
	},
	{
		id:   4,
		icon: "🍇",
	},
	{
		id:   5,
		icon: "🍒",
	},
	{
		id:   6,
		icon: "💎",
	},
	{
		id:   7,
		icon: "🍏",
	},
}

func NewSlotMachine() *SlotMachineType {
	return &SlotMachineType{
		Reels:   [][]SymbolType{},
		Credits: 50,
	}
}

func (s *SlotMachineType) Spin() {
	// if s.Credits < 10 return
	s.Reels = make([][]SymbolType, 3)
	for i := range s.Reels {
		s.Reels[i] = make([]SymbolType, 3)
		for j := range s.Reels[i] {
			s.Reels[i][j] = generateRandomSymbol(0, 7)
		}
	}
}

func (s *SlotMachineType) Display() {
	clearScreen()
	for i := range s.Reels {
		fmt.Print("\n          |")
		for j := range s.Reels[i] {
			fmt.Print(strings.TrimSpace(s.Reels[i][j].icon), " | ")
		}
	}
	fmt.Print("\n")

	// fmt.Println("--------------------------------------")
	// fmt.Printf("Credits: %d \n\n", s.Credits)
	// fmt.Println("Press (X) to play more ")
	// fmt.Println("--------------------------------------")
}

func (s *SlotMachineType) CheckWin() bool {
	for i := 0; i < 3; i++ {
		if s.Reels[i][0] == s.Reels[i][1] && s.Reels[i][1] == s.Reels[i][2] {
			return true
		}
		if s.Reels[0][i] == s.Reels[1][i] && s.Reels[1][i] == s.Reels[2][i] {
			return true
		}
	}

	//? DIAGONAL
	if s.Reels[0][0] == s.Reels[1][1] && s.Reels[1][1] == s.Reels[2][2] {
		return true
	}
	if s.Reels[0][2] == s.Reels[1][1] && s.Reels[1][1] == s.Reels[2][0] {
		return true
	}

	return false
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
