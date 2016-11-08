package tictactoe

type Player int

const (
	O Player = iota
	X
	EMPTY
)

type Board [3][3]Player // [row][col]

func (b *Board) deepCopyBoard() *Board {
	b2 := new(Board)
	for i, row := range(b) {
		for j, space := range(row) {
			b2[i][j] = space
		}
	}
	return b2
}

// Parse a Board from a string of 'o', 'x', and ' '
func FromString(param string) *Board {
	if len(param) != 9 {
		return nil
	}

	var output = new(Board)
	for i, c := range param {
		switch c {
		case ' ':
			output[i/3][i%3] = EMPTY
		case 'o':
			output[i/3][i%3] = O
		case 'x':
			output[i/3][i%3] = X
		default:
			return nil
		}
	}
	return output
}

// Print a Board out as nine characters - 'o', 'x', and ' '
func (b *Board) ToString() string {
	output := ""
	for _, row := range b {
		for _, space := range row {
			switch space {
			case EMPTY:
				output += " "
			case O:
				output += "o"
			case X:
				output += "x"
			}
		}
	}
	return output
}

// Check if it looks like its our turn
func (b *Board) Validate() string {
	xcount, ocount := 0, 0
	for _, row := range b {
		for _, space := range row {
			switch space {
			case O:
				ocount += 1
			case X:
				xcount += 1
			}
		}
	}

	// There are nine spaces on the board; make sure at least one is empty
	if xcount+ocount == 9 {
		return "All spaces are full; there's no where for me to play"
	}

	// Either the # of moves per player should be equal (we went first)
	// Or there should be one more X than O's (they went first)
	if xcount == ocount || xcount-1 == ocount {
		return ""
	}
	return "It's not O's turn."

}
