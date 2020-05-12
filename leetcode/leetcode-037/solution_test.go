package leetcode037

import (
	"testing"
	"fmt"
)

func TestSolution(t *testing.T) {
	fmt.Print("123")
	data := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(data)
}

// 添加约束, 使用dfs深度优先遍历进行迭代
func solveSudoku(board [][]byte) {
	backtrace(board, 0, 0)
	return
}

// row 从0开始, 达到9就可以退出了
func backtrace(board [][]byte, i int, j int) bool {
	// 横纵终止
	const m, n = 9, 9
	if j == n { // 当填满的时候就可以返回了
		return backtrace(board, i+1, 0)
	}
	if i == m {
		return true // 可行解, 那么直接终止了
	}
	// 不需要填数字
	if board[i][j] != '.' {
		ret := backtrace(board, i, j+1)
		fmt.Println(i, j)
		if j <= 9  {
			fmt.Println(i, j)
			for _, item := range board[0]{
				fmt.Print(string(item), " ")
			}
			fmt.Print("/n")
		}
		return ret 
	}

	for c := byte('1') ; c <= byte('9') ; c ++ {
		if isValid(board, i, j, c) { // 只需要判断横纵以及九宫格的情况即可也就是需要计算最多27个数
			// 设值并且前进, 并且向后推进
			board[i][j] = c
			if backtrace(board, i, j+1) { // 找到解了那就直接返回了不再继续了
				return true
			}
			board[i][j] = '.' // 恢复
		}
	}
	return false
}

func isValid(board [][]byte, i int, j int, num byte) bool {
	mark0 := 1 << (num - '1') // 九宫格
	mark1 := 1 << (num - '1') // 横
	mark2 := 1 << (num - '1') // 纵
	// 横
	for c := 0 ; c < 9; c ++ {
		if board[i][c] >= '1' && board[i][c] <= '9' {
			num := 1 << (board[i][c] - '1')
			if mark1 & num != 0 {
				return false
			}
			mark1 |= num
		}
	}
	// 纵
	for r := 0 ; r < 9; r ++ {
		if board[r][j] >= '1' && board[r][j] <= '9' {
			num := 1 << (board[r][j] - '1')
			if mark2 & num != 0 {
				return false
			}
			mark2 |= num
		}
	}
	// 九宫格
	rmin, rmax := i- i%3, i+3 - i%3
	cmin, cmax := j- j%3, j+3 - j%3
	for r := rmin; r < rmax ; r++{
		for c := cmin; c < cmax; c++{
			if board[r][c] >= '1' && board[r][c] <= '9' {
				num := 1 << (board[r][c] - '1')
				if mark0 & num != 0 {
					return false
				}
				mark0 |= num
			}
		}
	}
	return true
}
