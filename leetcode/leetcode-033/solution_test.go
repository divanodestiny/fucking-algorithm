package leetcode034

import (
	"testing"
)

func TestSolution(t *testing.T) {
	search([]int{5,7,7,8,8,10}, 8)
}

func search(nums []int, target int) int {
	lenght := len(nums)
	if lenght == 0 {
		return -1
	}
	l, r := 0, lenght - 1
	for l < r {
		m := l + (r - l) / 2 //  左中位数
		if nums[m] == target {
			return m
		} else if nums[r] == target {
			return r
		}
		if nums[l] <= nums[m] { // 左边有序，对左侧进行二分
			if nums[l] <= target && target <= nums[m] {
				r = m
			} else {
				l = m + 1
			}
		} else {
			if nums[m] < target && target <= nums[r] {
				l = m + 1
			} else {
				r = m
			}
		}
	}
	return -1
}