package main

func plusOne(digits []int) []int {
	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1]++
		return digits
	} else {
		if len(digits) == 1 {
			return []int{1, 0}
		}
		digits = append(plusOne(digits[:len(digits)-1]), 0)
	}
	return digits
}
