package main

import (
	"math/rand"
	"time"
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

func generateRandomId(min, max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	return rng.Intn(max-min) + 1
}
