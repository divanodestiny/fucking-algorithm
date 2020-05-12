package leetcode121

import (
	"math"
	"testing"
)

func TestSolution(t *testing.T) {
	maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4})
	maxProfit([]int{1, 2, 3, 4, 5})
	maxProfit([]int{7, 6, 4, 3, 1})
}

func TestMinInt(t *testing.T) {
	val := int(math.MinInt64)
	t.Log(val)
}

func max(args ...int) int {
	var m = args[0]
	for _, a := range args {
		if a > m {
			m = a
		}
	}
	return m
}

func min(args ...int) int {
	var m = args[0]
	for _, a := range args {
		if a < m {
			m = a
		}
	}
	return m
}

// 递推公式：
//dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
//dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
// 本质上只会设计dp[i]和dp[i-1]两个维度，所以申请空间只需要dp[2][k+1][2]就足够了空间复杂度是O(k)
// 每一次迭代应该只需要进行一次swap，将dp[0]与dp[1]交换即可
// 那么问题在于如何进行base case的设置
// base case的设置就是数学归纳法的前置条件
// 对于股票问题来说，前置条件就是dp[0]的基础值设置
// 存在以下情形：使用minInt来表示不存在的情况
// (1) k == 0， 有dp[0][k][0] = 0，dp[0][k][1] = minInt但其实后一个赋值是没有必要做的，因为不会被下一轮使用到
// (2) k >= 1， 有dp[0][k][0] = minInt, dp[0][k][1] = -prices[0]
// 其实第一天的情况就是这样；要么买入，此时净利润为-prices[0]；要么不买入, 此时净利润为0；其他的情形是不存在的；那么只需要基于这样的前提进行迭代就可以了
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var k int = 2
	maxTimes := min(k, len(prices)/2)
	// get space
	var dp [2][][2]int
	dp[0] = make([][2]int, maxTimes+1)
	dp[1] = make([][2]int, maxTimes+1)

	// base case
	dp[1][0][0], dp[1][0][1] = 0, int(math.MinInt32)
	for i := 1; i <= maxTimes; i++ {
		dp[1][i][0], dp[1][i][1] = 0, -prices[0]
	}
	
	// iteration
	for i := 1; i < len(prices); i++ {
		// days
		// swap
		dp[0], dp[1] = dp[1], dp[0]

		// process j == 0
		for j := 1; j <= maxTimes; j++ {
			// times
			dp[1][j][0] = max(dp[0][j][0], dp[0][j][1]+prices[i])
			dp[1][j][1] = max(dp[0][j][1], dp[0][j-1][0]-prices[i])
		}
	}

	// get max profit
	var ret int = dp[1][maxTimes][0]
	return ret
}
