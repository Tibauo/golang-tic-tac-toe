package morpion

//DisplayRule show rule
func DisplayRule(g *Game) {
	println("User 1 = ", g.J1, "\nUser 2 = ", g.J2)
	println(`
       j
     0 1 2
  0 |_|_|_|
i 1 |_|_|_|
  2 |_|_|_|
	`)
}

//DisplayBoard Show plate of the game
func DisplayBoard(g *Game) {
	println("       j   ")
	println("     0 1 2 ")
	for i := 0; i < 3; i++ {
		if i == 1 {
			print("i ", i, " |")
		} else {
			print("  ", i, " |")
		}
		for j := 0; j < 3; j++ {
			print(g.Board[i][j], "|")
		}
		println()
	}
	println("\n")
}
