package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	//按照区间起始位置对区间进行排序，以便后续合并操作
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//初始化合并结果切片，将第一个区间作为初始元素
	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		//如果当前区间的起始位置小于等于合并结果中最后一个区间的结束位置，说明存在重叠
		if intervals[i][0] <= merged[len(merged)-1][1] {
			//合并区间，取两个区间结束位置的最大值
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], intervals[i][1])
		} else {
			// 如果不存在重叠，则直接将当前区间添加到合并结果中
			merged = append(merged, intervals[i])
		}
	}
	return merged
}
