package leetcode034

import (
	"testing"
)

func TestSolution(t *testing.T) {
	search([]int{1,3,1,1,1}, 3)
}

// 有可能存在重复元素，不可以直接套用033的做法，需要有一定区别
func search(nums []int, target int) bool {
	lenght := len(nums)
	l, r := 0, lenght - 1
	for l <= r {
		m := l + (r - l) / 2 //  左中位数
		if nums[m] == target {
			return true
		} else if nums[r] == target {
			return true
		}
		if nums[l] == nums[m] {
			// 退化为顺序查找
			l++
		}else if nums[l] < nums[m] {   // 有重复元素的情况下没有办法通过nums[l] <= nums[m] 来判断左边有序还是右边有序
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
	return false
}