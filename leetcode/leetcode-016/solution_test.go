package leetcode16

import (
	"sort"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(threeSumClosest([]int{-1, 2, 1 , -4}, 1))
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return sum(nums)
	}
	sort.Ints(nums)
	ret := sum(nums[:3])
	for i := 0; i < len(nums); i++ {
		l, r := i+1, len(nums)-1
		for l < r {			
			s := nums[i] + nums[l] + nums[r]
			
			if abs(target - s) < abs(target - ret) {
				ret = s
			}

			if s > target {
				r --
			} else {
				l ++
			}
		}
	}
	return ret
}

func sum(nums []int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}

func abs(num int) int{
	if num > 0 {
		return num
	}
	return -1 * num
}
