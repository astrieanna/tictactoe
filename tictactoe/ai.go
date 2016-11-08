package tictactoe
import(
"log"
"fmt"
"strings"
)

type GameNode struct {
	b    *Board
	i, j int    // where are they playing
	p    Player // whose turn is it
}

func (b *Board) CheckForWin() *Player {
	var winners map[Player]bool
	winners = make(map[Player]bool)

	// Check for a winner
	// rows
	for _, row := range b {
		if row[0] == row[1] && row[1] == row[2] {
			winners[row[0]] = true
		}
	}

	// cols
	for j := 0; j < 3; j++ {
		if b[0][j] == b[1][j] && b[1][j] == b[2][j] {
			winners[b[0][j]] = true
		}
	}

	// diags
	// top left to bottom right;		top right to bottom left
	if (b[0][0] == b[1][1] && b[1][1] == b[2][2]) || (b[0][2] == b[1][1] && b[1][1] == b[2][0]) {
		winners[b[1][1]] = true
	}

	if winners[O] && winners[X] {
		log.Fatal("Multiple winners on board '" + b.ToString() + "'")
	} else if winners[O] {
		o := O
		return &o
	} else if winners[X] {
		x := X
		return &x
	}

	// Check for a draw
	filledIn := 0
	for _, row := range b {
		for _, space := range row {
			if space != EMPTY {
				filledIn++
			}
		}
	}
	if filledIn == 9 {
		e := EMPTY
		return &e
	}

	// Neither a winner nor a draw has happened yet
	return nil
}

func (b *Board) pickMoveHelper(p Player, i int, j int, depth int) (bestResult Player) {
	b[i][j] = p
	winner := b.CheckForWin()
	if winner != nil {
		return *winner
	}

	if p == X {
		bestResult = O
	} else {
		bestResult = X
	}
	for i, row := range b {
		for j, space := range row {
			if space == EMPTY {
				b2 := b.deepCopyBoard()
				var result Player // EMPTY == TIE
				if p == O {
					result = b2.pickMoveHelper(X, i, j, depth+1)
				} else {
					result = b2.pickMoveHelper(O, i, j, depth+1)
				}

				if result == p { // this move lets us win
					fmt.Printf(strings.Repeat(" ", depth) + "%d can win at %d, %d on %q\n", p, i, j, b.ToString())
					bestResult = result
				} else if result == EMPTY && bestResult != p { // this moves gets us to tie
					bestResult = result
				}
			}
		}
	}
	return
}

func (b *Board) pickMove() (bestR int, bestC int) {
	bestR, bestC = -1, -1
	bestResult := X
	for i, row := range b {
		for j, space := range row {
			if space == EMPTY {
				b2 := b.deepCopyBoard()
				fmt.Printf("Considering %d, %d on %q\n", i, j, b.ToString())
				var result Player // EMPTY == TIE
				result = b2.pickMoveHelper(O, i, j, 1)

				// we win or		we draw (and don't have a way to win)
				if result == O || (result == EMPTY && bestResult == X) {
					bestR, bestC, bestResult = i, j, result
				}
			}
		}
	}
	return
}

func (b *Board) MakeMove() {
	i, j := b.pickMove()
	b[i][j] = O // make the move we picked
}
