// O(nlogn)

import (
    "sort"
)

func minPairSum(nums []int) int {
    sort.Ints(nums)
    n := len(nums)
    max := 0

    for i := 0; i < n/2; i++ {
        sum := nums[i] + nums[n-i-1]
        if max < sum {
            max = sum
        }
    }

    return max
}
