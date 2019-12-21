package morpion

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//Game  Struct of Board Game
type Game struct {
	Board  [3][3]string
	Turn   int
	J1     string
	J2     string
	Player int
}

//New Create Game
func New() *Game {
	g := &Game{
		Board: [3][3]string{
			{"_", "_", "_"},
			{"_", "_", "_"},
			{"_", "_", "_"},
		},
		Turn:   0,
		J1:     "X",
		J2:     "O",
		Player: 1,
	}

	return g
}

// Player Show player
func Player(g *Game) {
	if g.Turn%2 == 0 {
		g.Player = 1
	} else {
		g.Player = 2
	}
}

// CheckMoove of player
func CheckMoove(g *Game, i int, j int) (bool, string) {
	if i > 2 {
		return true, "Value i is superior at 2"
	} else if j > 2 {
		return true, " Value j is superior at 2"
	} else if g.Board[i][j] != "_" {
		return true, "Case already use"
	} else if g.Turn%2 == 0 {
		g.Board[i][j] = "X"
		return false, ""
	} else {
		g.Board[i][j] = "O"
		return false, ""
	}

}

// Play write value of gamer in the good case
func Play(g *Game) {
	for {
		println("Player", g.Player, ": ")
		reader := bufio.NewReader(os.Stdin)
		print("i : ")
		texti, _ := reader.ReadString('\n')
		i, _ := strconv.Atoi(strings.Trim(texti, "\n"))
		print("j : ")

		textj, _ := reader.ReadString('\n')
		j, _ := strconv.Atoi(strings.Trim(textj, "\n"))
		check, error := CheckMoove(g, i, j)
		if check {
			println("error", error)
		} else {
			break
		}
	}
}

//Check the value
func Check(g *Game) bool {
	value := true

	if g.Board[1][1] != "_" && strings.Compare(g.Board[0][0], g.Board[1][1]) == 0 && strings.Compare(g.Board[1][1], g.Board[2][2]) == 0 {
		value = false
	} else if g.Board[1][1] != "_" && g.Board[2][0] == g.Board[1][1] && g.Board[1][1] == g.Board[0][2] {
		value = false
	} else if g.Board[0][0] != "_" && g.Board[0][0] == g.Board[0][1] && g.Board[0][1] == g.Board[0][2] {
		value = false
	} else if g.Board[1][0] != "_" && g.Board[1][0] == g.Board[1][1] && g.Board[1][1] == g.Board[1][2] {
		value = false
	} else if g.Board[2][0] != "_" && g.Board[2][0] == g.Board[2][1] && g.Board[2][1] == g.Board[2][2] {
		value = false
	} else if g.Board[0][0] != "_" && g.Board[0][0] == g.Board[1][0] && g.Board[1][0] == g.Board[2][0] {
		value = false
	} else if g.Board[0][1] != "_" && g.Board[0][1] == g.Board[1][1] && g.Board[1][1] == g.Board[2][1] {
		value = false
	} else if g.Board[0][2] != "_" && g.Board[0][2] == g.Board[1][2] && g.Board[1][2] == g.Board[2][2] {
		value = false
	}
	if value == false {
		println("You Win player : ", g.Player)
	}
	return value
}
