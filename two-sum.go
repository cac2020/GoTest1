package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		minus := target - num
		if index, ok := numMap[minus]; ok {
			return []int{i, index}
		} else {
			numMap[num] = i
		}
	}
	return []int{}
}
