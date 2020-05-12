package leetcode042

import (
	"testing"
)


func TestSolution(t *testing.T) {
	// trap([]int{0,1,0,2,1,0,1,3,2,1,2,1})
	// trap2([]int{0,1,0,2,1,0,1,3,2,1,2,1})
	trap3([]int{0,1,0,2,1,0,1,3,2,1,2,1})
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 动态规划
func trap(height []int) int {

	var ldp = make([]int, len(height))
	var rdp = make([]int, len(height))

	// 左边最高
	for i := 1; i < len(height) ; i++ {
		ldp[i] = max(ldp[i-1], height[i-1])
	}
	// 右边最高
	for i := len(height) - 2; i >= 0 ; i-- {
		rdp[i] = max(rdp[i+1], height[i+1])
	}

	res := 0 
	for i := 1 ; i < len(height) - 1 ; i ++ {
		res += max(0, min(ldp[i], rdp[i]) - height[i])
	}
	return res
}

type Stack []int

func (s *Stack) Push(i int) {
	*s = append(*s, i)
}

func (s Stack)Top()int {
	return s[len(s)-1]
}

func (s *Stack) Pop() int {
	old := *s
	ret := s.Top()
	*s = old[:len(old)-1]
	return ret
}

func (s Stack) Empty() bool {
	return len(s) == 0
}


// 基于栈
func trap2(height []int)int {

	var s = make(Stack, 0, len(height)) // 使用数组来模拟栈
	res := 0

	for i := 0 ; i < len(height) ; i ++ {
		for !s.Empty() && height[s.Top()] < height[i] {
			top := s.Pop()
			if s.Empty() {
				break
			}
			w := i - s.Top() - 1                                // 注意宽度计算方式
			h := min(height[s.Top()], height[i]) - height[top]  // 注意高度计算方式
			res += w * h

		}
		s.Push(i)
	}

	return res
}

// 双指针
// 双指针的理解是基于dp的搞法进行空间复杂度和时间复杂度的优化，因为ml[i]和mr[i]其实都只用了一次
// 1. 初步优化思路
// ml[i] = max(ml[i-1], height[i-1])
// mr[i] = max(ml[i+1], height[i+1])
// 是可以裁剪掉一个，先从右到左计算出mr数组，然后在从左到右推进的过程中，一边计算ml，一边计算柱的大小
// 那么这样也仅仅是少了一轮循环，时间复杂度和空间复杂度是O(n), 如何考虑把mr也裁剪掉就是关键了
// 2. 递进优化思路
// 每一个高度的计算方式是min(ml[i], mr[i]) = min(max(ml[i-1], height[i-1]), ml[i+1], height[i+1])
// 而初始边界ml[0], mr[max]都为0
// 那么对于从左边递推的数来说， 要保证ml是边界的话，只要左边界指向的数小于等于右边界即可
// 对于从右边递推的数来说，要保证mr是边界，需要右边界指向的数大于等于左边界
func trap3(height []int) int {

	ml, mr := 0, 0              // 左右边界
	l, r := 1, len(height) - 2  // 左右指针
	ret := 0
	for l <= r {
		if height[l-1] <= height[r+1] { // 左右边界比较
			ml = max(ml, height[l-1])
			if height[l] < ml {
				ret += ml - height[l]
			}
			l ++
		} else {
			mr = max(mr, height[r+1])
			if height[r] < mr {
				ret += mr - height[r]
			}
			r --
		}
	}
	return ret
}