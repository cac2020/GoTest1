package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(prefix) {
			//如果小于前缀长度 则前缀按照当前字符串长度截取并赋值
			prefix = prefix[:len(strs[i])]
		}
		for j := 0; j < len(prefix); j++ {
			if strs[i][j] != prefix[j] {
				prefix = prefix[:j]
			}
		}
	}
	return prefix
}
