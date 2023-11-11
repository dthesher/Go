package dslice

import (
	"math"
	"reflect"
)

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
//// Make sure to pass second param: element with type casting if it is not "int"
//// Returns -1 for index if it does not find element in slice
func Contains(slice interface{}, element interface{}) int {
	slc := InterfaceSlice(slice)
	if slc != nil {
		for i, a := range slc {
			if a == element {
				return i
			}
		}
	}
	return -1
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

func MinFromArray(arr []int) int {
	if len(arr) == 0 {
		panic("Empty array")
	}

	min := math.MaxInt
	for _, value := range arr {
		if value < min {
			min = value
		}
	}

	return min
}
