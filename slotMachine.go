package main

import (
	"math/rand"
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
	s.Reels = make([][]SymbolType, 3)
	for i := range s.Reels {
		s.Reels[i] = make([]SymbolType, 3)
		for j := range s.Reels[i] {
			s.Reels[i][j] = generateRandomSymbol(1, 8)
		}
	}
}

func generateRandomSymbol(min, max int) SymbolType {
	//? 	THIS IS BASED ON TIME - GO TOO FAST IT GIVES SAME NUMBER ID
	// source := rand.NewSource(time.Now().UnixNano())
	// rng := rand.New(source)

	return Symbols[rand.Intn(max-min)+1]
}
