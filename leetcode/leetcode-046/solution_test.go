package leetcode046

import (
	"testing"
)


// 最大堆
type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] > h[j] } // 最大堆用>, 最小堆用<
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
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

func TestSolution(t *testing.T){
	permute([]int{1,2,3})
}

// dfs
func permute(nums []int) [][]int {
	var ret [][]int
	var backtrace func(path []int, options []int)
	backtrace = func(path []int, options []int) {
		if len(options) == 0 {
			newPath := make([]int, len(path))
			copy(newPath, path)
			ret = append(ret, newPath)
			return
		}
		for k, v := range options {
			// 选择
			newOption := make([]int,0, len(options)-1)
			newOption = append(append(newOption, options[:k]...), options[k+1:]...)
			backtrace(append(path, v), newOption)
		}
	}
	backtrace(nil, nums)
	return ret
}
