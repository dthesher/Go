package dSlice

import "reflect"

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
* @params "slice" slice identifier,
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

/**
* convert interface{} (which have slice in it) to to []interface{}.
* @params "slice" slice identifier,
* @return []interface{} or nil. */
//// It will panic if you give to a non slice type interface
//// if it have empty slice then it will return nil
func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}
