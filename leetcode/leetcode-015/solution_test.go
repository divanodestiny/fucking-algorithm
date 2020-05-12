package leetcode15

import (
	"fmt"
	"sort"
	"testing"
)

func Test(t *testing.T) {
	// fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	// fmt.Println(threeSum([]int{0, 0, 0, 0}))
	fmt.Println(threeSum([]int{1,1,-2}))
}

func threeSum(nums []int) [][]int {
	var ret [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		// 选中nums[i]，然后用双指针来匹配另外两个数
		if nums[i] > 0 {
			break
		}

		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 双指针
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			
			if sum == 0 {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
			}
			
			if sum > 0 {
				// too small move r
				for { 
					if r--; nums[r] != nums[r+1] || l >= r { break }
				}
			} else {
				for {
					if l++; nums[l-1] != nums[l] || l >= r { break }
				}
			}
		}

	}
	return ret
}
