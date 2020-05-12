package leetcode462

import (
	"sort"
	"testing"
)


func TestSolution(t *testing.T){
	minMoves2([]int{1,2,3})
}

func minMoves2(nums []int) int {
	sort.Ints(nums)
	
	medium := (nums[(len(nums) -1)/2] + nums[len(nums) / 2]) / 2
	sum := 0
	for _, item := range  nums {
		if item > medium {
			sum += item - medium
		} else {
			sum += medium - item
		}
	}
	return sum
}