package sudoku


func isValidSudoku(board [][]byte) bool {

	boardMap := make(map[int][]byte)

	for k, row := range board {
		for _, v := range row {
			row := make(map[byte]int)
			if v == byte('.') {
				continue
			}

			ok := boardMap[k][v]
	
			if ok {
				return false
			}
	
			if row[v] {
				return false
			}
	
			row[v] = 1
			boardMap[k] = append(boardMap[k], v)
		} 
	}
	return true
}
