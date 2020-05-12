package leetcode034

import (
	"testing"
	"container/heap"
)

func TestSolution(t *testing.T) {
	hIndex([]int{0,1,3,5,6})
	hIndex([]int{0})
	hIndex([]int{0, 1})
}

func hIndex(citations []int) int {
	l, r := 0, len(citations) - 1
	if len(citations) == 0 || citations[r] == 0 {
		return 0
	}
	for l < r {
		m := l + (r - l)/2 // left medium
		more := len(citations) - m // 大于等于m的数量
		if more > citations[m] { // 大于的情况，说明m取值偏小， 可以向后二分向后取二分
			l = m + 1
		} else  { // 小于的情况， 说明m取的过大，移动右边
			r = m
		}
	}
	return len(citations) - l
}