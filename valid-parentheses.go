package main

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}

	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			if (s[i] == ')' && stack[len(stack)-1] == '(') || (s[i] == ']' && stack[len(stack)-1] == '[') || (s[i] == '}' && stack[len(stack)-1] == '{') {
				stack = stack[:len(stack)-1]
			}
		}
	}

	return len(stack) == 0
}
