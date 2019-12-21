package main

import (
	"morpion"
)

func main() {
	g := morpion.New()
	println("Game")
	var game bool = true
	morpion.DisplayRule(g)
	for game {
		morpion.Player(g)
		morpion.Play(g)
		g.Turn++
		game = morpion.Check(g)
		morpion.DisplayBoard(g)
	}
}
