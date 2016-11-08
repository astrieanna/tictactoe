package tictactoe

func (b *Board) MakeMove() {
	for i, row := range b {
		for j, space := range row {
			if space == EMPTY {
				b[i][j] = O
				return
			}
		}
	}
}
