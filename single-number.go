package main

func singleNumber(nums []int) int {
	mp := make(map[int]int)
	for _, item := range nums {
		if _, exists := mp[item]; exists {
			mp[item]++
		} else {
			mp[item] = 1
		}
	}

	for key, val := range mp {
		if val == 1 {
			return key
		}
	}
	return 0
}
