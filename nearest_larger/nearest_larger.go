package nearest_larger

import (
	"math"
)

// [1, 2], 0
func checkIndexes(nums []int, checkValue, startIndex, endIndex int) int {
	// loops bet 2 values
	// checks currentValue against checkValue
	// return if currenValue > checkValue
	// else return -1
	for i := startIndex; i < endIndex; i++ {
		if nums[i] > checkValue {
			return i
		}
	}

	return -1
}

func NearestLarger(nums []int, i int) int {
	leftIndex := checkIndexes(nums, nums[i], 0, i)
	leftDistance := math.Abs(float64(leftIndex - i))

	rightIndex := checkIndexes(nums, nums[i], i, len(nums))
	rightDistance := math.Abs(float64(rightIndex - i))

	if leftIndex != -1 && leftDistance <= rightDistance {
		return leftIndex
	}

	return rightIndex
}
