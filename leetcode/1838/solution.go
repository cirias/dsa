// O(nlogn)

import (
	"sort"
)

func maxFrequency(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)

	maxFreq := 1

	// maintain the most frequent sequence as the range of [i, j]
	i, j := 0, 0
	for j+1 < n {
		if k >= 0 {
			// add one number into the frequent sequence if we still have remaining operations
			j += 1
			k -= (nums[j] - nums[j-1]) * (j - i)
		} else {
			// or drop one if there is no remaining operations
			k += nums[j] - nums[i]
			i += 1
		}

		if k >= 0 {
			// whenever this is a valid operation, update the maxFreq to the current length of the sequence
			maxFreq = max(maxFreq, j-i+1)
		}
	}

	// continue dropping until the sequence become possible again
	for k < 0 {
		k += nums[j] - nums[i]
		i += 1
	}
	maxFreq = max(maxFreq, j-i+1)

	return maxFreq
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
