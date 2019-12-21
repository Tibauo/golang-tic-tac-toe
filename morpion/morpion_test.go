package morpion

import (
	"testing"
)

func TestPlayerJ1(t *testing.T) {
	g := New()
	g.Turn = 10
	if g.Turn%2 == 0 {
		g.Player = 1
	} else {
		g.Player = 2
	}
	if g.Player == 2 {
		t.Errorf("TestPlayerJ1")
	}
}
func TestPlayerJ2(t *testing.T) {
	g := New()
	g.Turn = 9
	if g.Turn%2 == 0 {
		g.Player = 1
	} else {
		g.Player = 2
	}
	if g.Player == 1 {
		t.Errorf("TestPlayerJ2")
	}
}
func TestSucessMoove(t *testing.T) {
	g := New()
	test, _ := CheckMoove(g, 1, 2)
	if test {
		t.Errorf("error: TestSucessMoove")
	}
}

func TestFailedMoovei(t *testing.T) {
	g := New()
	test, _ := CheckMoove(g, 5, 2)
	if !test {
		t.Errorf("error: TestFailedMoovei")
	}
}

func TestFailedMoovej(t *testing.T) {
	g := New()
	test, _ := CheckMoove(g, 1, 4)
	if !test {
		t.Errorf("error: TestFailedMoovej")
	}
}
func TestCheckDiagSuccess(t *testing.T) {
	g := New()
	for i := 0; i < 3; i++ {
		g.Board[i][i] = "X"
	}
	checkDiag := Check(g)
	if checkDiag {
		t.Errorf("error: TestCheckDiagSuccess")
	}
}

func TestCheckDiagNotComplete(t *testing.T) {
	g := New()
	for i := 0; i < 2; i++ {
		g.Board[i][i] = "X"
	}
	checkDiag := Check(g)
	if !checkDiag {
		t.Errorf("error: TestCheckDiagNotComplete")
	}
}

func TestCheckDiagJ1andJ2(t *testing.T) {
	g := New()
	g.Board[1][1] = "X"
	g.Board[0][0] = "0"
	g.Board[2][2] = "X"
	checkDiag := Check(g)
	if !checkDiag {
		t.Errorf("error: TestCheckDiagJ1andJ2")
	}
}

func TestCheckSuccessColumn0(t *testing.T) {
	g := New()
	for i := 0; i < 3; i++ {
		g.Board[i][0] = "X"
	}
	checkColumn := Check(g)
	if checkColumn {
		t.Errorf("error: TestCheckSuccessColumn0")
	}
}

func TestCheckSuccessColumn1(t *testing.T) {
	g := New()
	for i := 0; i < 3; i++ {
		g.Board[i][1] = "X"
	}
	checkColumn := Check(g)
	if checkColumn {
		t.Errorf("error: TestCheckSuccessColumn1")
	}
}

func TestCheckSuccessColumn2(t *testing.T) {
	g := New()
	for i := 0; i < 3; i++ {
		g.Board[i][2] = "X"
	}
	checkColumn := Check(g)
	if checkColumn {
		t.Errorf("error: TestCheckSuccessColumn2")
	}
}

func TestCheckFailedColumn0(t *testing.T) {
	g := New()
	for i := 0; i < 2; i++ {
		g.Board[i][0] = "X"
	}
	checkColumn := Check(g)
	if !checkColumn {
		t.Errorf("error: TestCheckFailedColumn0")
	}
}

func TestCheckFailedColumn1(t *testing.T) {
	g := New()
	for i := 0; i < 2; i++ {
		g.Board[i][1] = "X"
	}
	checkColumn := Check(g)
	if !checkColumn {
		t.Errorf("error: TestCheckFailedColumn1")
	}
}

func TestCheckFailedColumn2(t *testing.T) {
	g := New()
	for i := 0; i < 2; i++ {
		g.Board[i][2] = "X"
	}
	checkColumn := Check(g)
	if !checkColumn {
		t.Errorf("error: TestCheckFailedColumn2")
	}
}

func TestCheckSuccessLine0(t *testing.T) {
	g := New()
	for j := 0; j < 3; j++ {
		g.Board[0][j] = "X"
	}
	checkColumn := Check(g)
	if checkColumn {
		t.Errorf("error: TestCheckSuccessLine0")
	}
}

func TestCheckSuccessLine1(t *testing.T) {
	g := New()
	for j := 0; j < 3; j++ {
		g.Board[1][j] = "X"
	}
	checkColumn := Check(g)
	if checkColumn {
		t.Errorf("error: TestCheckSuccessLine1")
	}
}

func TestCheckSuccessLine2(t *testing.T) {
	g := New()
	for j := 0; j < 3; j++ {
		g.Board[2][j] = "X"
	}
	checkColumn := Check(g)
	if checkColumn {
		t.Errorf("error: TestCheckSuccessLine2")
	}
}

func TestCheckFailedLine0(t *testing.T) {
	g := New()
	for j := 0; j < 2; j++ {
		g.Board[0][j] = "X"
	}
	checkColumn := Check(g)
	if !checkColumn {
		t.Errorf("error: TestCheckFailedLine0")
	}
}

func TestCheckFailedLine1(t *testing.T) {
	g := New()
	for j := 0; j < 2; j++ {
		g.Board[1][j] = "X"
	}
	checkColumn := Check(g)
	if !checkColumn {
		t.Errorf("error: TestCheckFailedLine1")
	}
}

func TestCheckFailedLine2(t *testing.T) {
	g := New()
	for j := 0; j < 2; j++ {
		g.Board[2][j] = "X"
	}
	checkColumn := Check(g)
	if !checkColumn {
		t.Errorf("error: TestCheckFailedLine2")
	}
}
