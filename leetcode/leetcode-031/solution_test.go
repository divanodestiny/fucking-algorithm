package leetcode031

import (
	"testing"
)

func TestSolution(t *testing.T) {
	// nextPermutation([]int{1,2,3})
	nextPermutation([]int{3,2,1})
	// nextPermutation([]int{1,1,5})
	// nextPermutation([]int{2,3,1})
}

func nextPermutation(nums []int)  {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {// 从后往前遍历,找到第一个峰
		i--
	}

	if i >= 0 {
		j := len(nums) - 1 // 找到最后一个大于nums[i]的数字nums[j]
		for nums[j] <= nums[i] {
			j --
		}
		nums[i], nums[j] = nums[j], nums[i]// 交换
	}


	// 反转nums[i+1:len(nums)] 
	l, r := i+1, len(nums)-1
	for l < r{
		nums[l], nums[r] = nums[r], nums[l]
		l ++
		r --
	}
	return 
}