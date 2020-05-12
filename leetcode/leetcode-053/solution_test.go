package leetcode053

import (
	"testing"
)

func TestSolution(t *testing.T){
	maxSubArray2([]int{-2,1,-3,4,-1,2,1,-5,4})
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	// dp 算法
	// 递推公式 dp[i] = max(dp[i-1] + nums[i], nums[i])
	if len(nums) == 0 {
		return 0
	}
	ret := nums[0]
	prev := nums[0]
	for i := 1; i < len(nums); i++{
		n := nums[i]
		prev = max(prev+n, n) // 简单来讲一边迭代一边放弃掉小于0的段落
		ret = max(ret, prev)
	}
	return ret
}


func maxSubArray2(nums []int) int {
	// 分治法
	_, _, ret , _ := divideAndConquer(nums)
	return ret
}

func divideAndConquer(nums []int)(lmax int, rmax int, tmax int, sum int) {
	if len(nums) == 1 {
		return nums[0], nums[0],nums[0],nums[0]
	}
	// 分成两个部分
	m := len(nums) / 2
	
	leftPart := nums[:m]
	rightPart := nums[m:]
	
	l0, l1, l2, l3 := divideAndConquer(leftPart)
	r0, r1, r2, r3 := divideAndConquer(rightPart)
	
	lmax = max(l0, l3+r0)
	rmax = max(r1, r3+l1)
	tmax = max(l1+r0, max(l2, r2))
	sum = l3 + r3
	return 
}


