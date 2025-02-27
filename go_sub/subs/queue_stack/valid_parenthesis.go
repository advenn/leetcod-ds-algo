package queue_stack

func IsValid(s string) bool {
	var stack []string
	for index := 0; index < len(s); index++ {
		b := string(s[index])
		if b == "(" || b == "[" || b == "{" {
			stack = append(stack, b)
		} else {
			if len(stack) == 0 {
				return false
			}
			if b == ")" && stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
			} else if b == "]" && stack[len(stack)-1] == "[" {
				stack = stack[:len(stack)-1]
			} else if b == "}" && stack[len(stack)-1] == "{" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
