package leetcode121

import (
	"testing"
)

func TestSolution(t *testing.T) {
	maxProfit([]int{3,3,5,0,0,3,1,4})
	maxProfit([]int{1,2,3,4,5})
	maxProfit([]int{7,6,4,3,1})
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
//dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
func maxProfit(prices []int) int {

	var n = len(prices)
	if n == 0 {
		 return 0
	}
	var dp [][2]int = make([][2]int, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0], dp[i][1] = 0, -prices[0]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[n-1][0]
