package dslices

func MaxInts(nums []int) int {
	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

/**
* Check for the element is available or not in slice.
* @param "slice" slice identifier,
*        "element" the num or string you wish to check.
* @return this Student's name.
 */
func Contains(slice []interface{}, element interface{}) bool {
	for _, a := range slice {
		if a == element {
			return true
		}
	}
	return false
}
