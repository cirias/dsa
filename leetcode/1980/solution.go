// https://leetcode.com/problems/find-unique-binary-string/
// O(n)
func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	ret := make([]byte, n)

	for i := 0; i < n; i++ {
		// for each bit position i, the bit is flipped of the same bit position of nums[i]
		// therefore it is guaranteed to be different from all nums
		if nums[i][i] == '1' {
			ret[i] = '0'
		} else {
			ret[i] = '1'
		}
	}

	return string(ret)
}
