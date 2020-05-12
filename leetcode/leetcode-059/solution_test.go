package leetcode058

import (
	"testing"
)
func TestSolution(t *testing.T){
	generateMatrix(3)
	generateMatrix(4)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func generateMatrix(n int) [][]int {
	iMin, iMax := 0, n - 1
	jMin, jMax := 0, n - 1
	counter := 1
	ret := make([][]int, n)
	for i := 0; i < n ; i++{
		ret[i] =make([]int, n)
	}

	for iMin <= iMax && jMin <= jMax {
		for j := jMin; j <= jMax ; j ++{ // right
			ret[iMin][j] = counter
			counter ++
		}

		for i := iMin + 1; i <=iMax ; i ++{ // down
			ret[i][jMax] = counter
			counter ++
		}

		for j := jMax-1; j >= jMin ; j --{ // left
			ret[iMax][j] = counter
			counter ++
		}

		for i := iMax - 1; i > iMin ; i --{ // up
			ret[i][jMin] = counter
			counter ++
		}
		iMax --
		jMax --
		iMin ++
		jMin ++
	}



	return ret
}