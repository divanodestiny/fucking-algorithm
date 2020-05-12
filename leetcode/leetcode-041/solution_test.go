package leetcode041

import (
	"testing"
)

func TestSolution(t *testing.T){
	firstMissingPositive([]int{1,2,0})
	firstMissingPositive([]int{3,4,-1,1})
	firstMissingPositive([]int{7,8,9,11,12})
}

func firstMissingPositive(nums []int) int {

	// 第一轮 检查1是否存在
	hasOne := 0
	for _, n := range nums {
		if n == 1 {
			hasOne++
		}
	}
	if hasOne == 0 {
		return 1
	}

	// 第二轮 将非正数和超过长度的置位1
	for i, n := range nums{
		if n <= 0 || n > len(nums) {
			nums[i] = 1
		}
	}

	// 第三轮 标记
	for _, n := range nums {
		n = abs(n)
		if nums[n-1] > 0 {
			nums[n-1] *= -1
		}	
	}
	// 第四轮 拿到第一个正数
	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
} 

