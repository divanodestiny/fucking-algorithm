package leetcode300

import (
	"testing"
)

func TestSolution(t *testing.T){
	lengthOfLIS([]int{10,9,2,5,3,7,101,18})
}


func max(args ...int) int {
	max := args[0]
	l := len(args)
	for i := 0 ; i < l ; i ++{  
		if max < args[i]{
			max = args[i]
		}
	} 
	return max
}

// 遍历式方法 O(n^2)
func lengthOfLIS(nums []int) int {
	l := len(nums)
	dp := make([]int,l)
	ret := 0
	for i := 0 ; i < l ; i ++ {
		dp[i] = 1
		for j := 0 ; j < i ; j ++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
		if dp[i] > ret {
			ret = dp[i]
		}
	}
	return ret
}

// 贪心+二分 O(nlogn)
func lengthOfLIS2(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	dp := make([]int, 1, length)
	dp[0] = nums[0]
	for i := 1 ; i < length ; i ++ {
		// 二分
		l, r := 0 , len(dp) - 1
		if nums[i] > dp[r] { // 上界检查
			dp = append(dp, nums[i])
			continue
		}

		for l < r { // 找到地方存放
			m := (l + r) / 2 // 取m为左中位数 二分法看leetcode-30
            if dp[m] < nums[i] {
				l = m+1
			} else {
				r = m
			}
		}
		dp[l] = nums[i]		
	}
	return len(dp)
}