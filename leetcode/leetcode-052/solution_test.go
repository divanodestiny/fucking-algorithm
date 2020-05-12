package leetcode051

import (
	"testing"
)

func TestSolution(t *testing.T){
	totalNQueens(4)
}




func totalNQueens(n int) int{
	var res int
	// 构造棋盘
	board := make([][]byte, n)
	for i := 0 ; i < n ; i++{
		board[i] = make([]byte, n)
		for j := 0 ; j < n ; j ++ {
			board[i][j] = '.'
		}
	}

	var backtrace func(board [][]byte, i int) 
	backtrace = func(board [][]byte, i int) {
		if i == n {
			res ++
			return 
		}

		for j := 0 ; j < n ; j ++ {
			if isValid(board, i, j) {
				board[i][j] = 'Q'       // 设值
				backtrace(board, i + 1) // dfs
				board[i][j] = '.'       // 恢复
			}
		}
	}
	backtrace(board, 0)

	return res
}

func isValid(board [][]byte, i int, j int) bool {
	if i == 3 && j == 2{
		i ++
		i --
	}
	// 横
	for _, c := range board[i] {
		if c == 'Q'{
			return false
		}
	}
	// 纵
	for r := 0 ; r < len(board); r ++{
		if board[r][j] == 'Q' {
			return false
		}
	}
	// 对角1
	for r := 0 ; r < len(board); r ++ {
		c := r - i + j
		if c >= 0 && c < len(board) && board[r][c] == 'Q' {
			return false
		}
	}
	// 对角2
	for r := 0 ; r < len(board); r ++ {
		c := i + j - r
		if c >= 0 && c < len(board) && board[r][c] == 'Q' {
			return false
		}
	}

	// 有效
	return true
}

