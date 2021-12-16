package dMath

/* Pow() will return power on num */
func Pow(num int, power int) int {
	if power == 0 {
		return 1
	}

	if power%2 == 0 {
		return Pow(num*num, power/2)
	}

	return num * Pow(num, power-1)
}
