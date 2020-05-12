package leetcode034

import (
	"testing"
)

func TestSolution(t *testing.T) {
	searchRange([]int{5,7,7,8,8,10}, 8)
}

func searchRange(nums []int, target int) []int {
	lenght := len(nums)
	if lenght == 0 {
		return []int{-1, -1}
	}
	// 找到第一个位置
	l, r := 0 , lenght - 1
	for l < r {
		m := l + (r - l )/ 2 // 左中位数
		if nums[m] < target {
			l = m + 1
		} else {
			r = m 
		}
	}
	if nums[l] != target {
		return []int{-1, -1}
	}
	
	// 找到最后一个位置
	l1, r1 := l , lenght - 1
	for l1 < r1 {
		m := l1 + (r1 - l1 + 1)/2 // 右中位数（使用左中位数会造成死循环）
		if nums[m] > target {
			r1 = m - 1
		} else {
			l1 = m
		}
	}
	return []int{l, r1}
}