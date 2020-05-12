package leetcode034

import (
	"testing"
)

func TestSolution(t *testing.T) {
	findMin([]int{0,1,2,4,5,6,7})
}

func findMin(nums []int) int {
	lenght := len(nums)
	if lenght == 0 {
		return -1
	}
	// 找到第一个位置
	l, r := 0 , lenght - 1
	for l <= r { // 
		if nums[l] <= nums[r] { // 说明有序
			return nums[l]
		}
		m := l + (r - l )/ 2 // 左中位数
		if nums[l] <= nums[m] { // 左侧有序， 那么往右侧找
		     l = m + 1
		} else { // 右侧有序，那么往左侧找
			 r = m
		}
	}
	return -1
}