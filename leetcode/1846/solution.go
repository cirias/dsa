// https://leetcode.com/problems/maximum-element-after-decreasing-and-rearranging/
//
// The problem can be easily solved by O(nlogn) time complexity.
// Below is an overcomplicated solution with O(n) + O(mlogm) time complexity where m is the unique number of the array.

import (
	"sort"
)

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	// count the number of each number
	numCounts := make(map[int]int)
	for _, num := range arr {
		numCounts[num] += 1
	}

	// sort the unique numbers
	uniqNums := make([]int, 0, len(numCounts))
	for num := range numCounts {
		uniqNums = append(uniqNums, letter)
	}
	sort.Ints(uniqNums)

	maxium := 0
	for _, num := range uniqNums {
		// for each unique number, we advance the maxium by the minimum of the count of the number and the difference between the number and the current maximum
		maxium = maxium + min(counts[num], num-maxium)
	}
	return maxium
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
