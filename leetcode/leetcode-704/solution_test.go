package leetcode704

func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	l, r := 0, length - 1
	for l < r {
		m := l + (r - l) / 2
		if nums[m] < target {
			l = m + 1
		} else {
			r = m 
		}
	}

	if nums[l] == target {
		return l
	}
	
	return -1
}