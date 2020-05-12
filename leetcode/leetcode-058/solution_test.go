package leetcode058

import (
	"testing"
)
func TestSolution(t *testing.T){
	lengthOfLastWord("a ")
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

func lengthOfLastWord(s string) int {
	var ret int
	var flag = false
	for _, c := range s {
		if c == ' ' {
			flag = true
		} else if flag == true {
			ret = 1
			flag = false 
		} else {
			ret ++
		}
	}
	return ret
}