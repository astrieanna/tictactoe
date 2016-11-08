package tictactoe

type Player int

const (
	O Player = iota
	X
	EMPTY
)

type Board [3][3]Player  // [row][col]


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

func (b *Board) ToString() string {
	output := ""
	for _, row := range(b) {
		for _, space := range(row) {
		switch(space) {
			case EMPTY: output += " "
			case O: output += "o"
			case X: output += "x"
		}
	}
	}
	return output
}

func (*Board) validate() bool {
	return true
}
