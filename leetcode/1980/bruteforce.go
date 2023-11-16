// O(n^2)
import (
	"strconv"
	"fmt"
)

func findDifferentBinaryString(nums []string) string {
	taken := make(map[uint16]struct{})
	for _, binaryNum := range nums {
		num, _ := strconv.ParseUint(binaryNum, 2, 16)
		taken[uint16(num)] = struct{}{}
	}

	// format the binary to be length of n
	format := fmt.Sprintf("%%0%db", len(nums))

	// this is O(n) because there are at most n numbers in the map
	for i := 0; ; i++ {
		if _, ok := taken[uint16(i)]; !ok {
			return fmt.Sprintf(format, i)
		}
	}

	panic("not found")
}
