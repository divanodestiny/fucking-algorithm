package leetcode055

import (
	"testing"
)


func TestSolution(t *testing.T){
	canJump([]int{3,2,1,0,4})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canJump(nums []int) bool {
	maxIdx := 0
	for idx, num := range nums {
		if idx > maxIdx {
			break
		}
		m := idx + num
		if m > maxIdx {
			maxIdx = m
		}
	}
	if maxIdx >= len(nums) - 1 {
		return true
	}
	return false
}