package main

import "fmt"

func main() {
	//1.只出现一次的数字
	int1 := []int{2, 2, 1}
	fmt.Println("1.只出现一次的数字:", singleNumber(int1))
	int2 := []int{4, 1, 2, 1, 2}
	fmt.Println("1.只出现一次的数字:", singleNumber(int2))
	int3 := []int{1}
	fmt.Println("1.只出现一次的数字:", singleNumber(int3))

	//2.回文数
	fmt.Println("2.回文数:", isPalindrome(121))
	fmt.Println("2.回文数:", isPalindrome(-121))
	fmt.Println("2.回文数:", isPalindrome(10))

	//3. 有效的括号
	fmt.Println("3. 有效的括号:", isValid("()"))
	fmt.Println("3. 有效的括号:", isValid("()[]{}"))
	fmt.Println("3. 有效的括号:", isValid("(]"))
	fmt.Println("3. 有效的括号:", isValid("([])"))
	fmt.Println("3. 有效的括号:", isValid("([)]"))

	//4.最长公共前缀
	fmt.Println("4.最长公共前缀:", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println("4.最长公共前缀:", longestCommonPrefix([]string{"dog", "racecar", "car"}))

	//5.加一
	fmt.Println("5.加一:", plusOne([]int{1, 2, 3}))
	fmt.Println("5.加一:", plusOne([]int{4, 3, 2, 1}))
	fmt.Println("5.加一:", plusOne([]int{9}))
	fmt.Println("5.加一:", plusOne([]int{9, 9}))
	fmt.Println("5.加一:", plusOne([]int{9, 9, 9}))

	//6.删除有序数组中的重复项
	fmt.Println("6.删除有序数组中的重复项:", removeDuplicates([]int{1, 1, 2}))
	fmt.Println("6.删除有序数组中的重复项:", removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

	//7.合并区间
	fmt.Println("7.合并区间:", merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println("7.合并区间:", merge([][]int{{1, 4}, {4, 5}}))
	fmt.Println("7.合并区间:", merge([][]int{{4, 7}, {1, 4}}))

	//8.两数之和
	fmt.Println("8.两数之和:", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println("8.两数之和:", twoSum([]int{3, 2, 4}, 6))
	fmt.Println("8.两数之和:", twoSum([]int{3, 3}, 6))
}
