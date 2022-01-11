package number

import "math"

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func MaxFromMultiple(nums ...int) int {
	max := math.MinInt
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func MinFromMultiple(nums ...int) int {
	min := math.MaxInt
	for _, num := range nums {
		if num < min {
			num = min
		}
	}
	return min
}

func Sort(nums ...int) []int {
	if len(nums) == 1 {
		return nums
	} // 如果num数组只有1个数字，则返回它本身
	return bubbleSort(nums)
}

// Sort 冒泡排序
func bubbleSort(nums []int) []int {
	length := len(nums)
	flag := true // 优化
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			// 如果前者大于后者则交换位置（从小到大排序）
			if nums[j] > nums[j+1] {
				flag = false // 优化
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
		// 优化
		if flag {
			break
		}
	}
	return nums
}
