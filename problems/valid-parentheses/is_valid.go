package validparentheses

func IsValid(s string) bool { //nolint:varnamelen
	parentheses := map[int32]int32{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var stack []int32

	for _, char := range s {
		if value, ok := parentheses[char]; ok {
			stack = append(stack, value)

			continue
		}

		if len(stack) == 0 {
			return false
		}

		if char != stack[len(stack)-1] {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
