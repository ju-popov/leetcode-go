package romantointeger

func RomanToInt(s string) int { //nolint:cyclop,varnamelen
	result := 0

	for i := 0; i < len(s); i++ { //nolint:varnamelen
		current := s[i]

		var next uint8
		if i < len(s)-1 {
			next = s[i+1]
		}

		switch {
		case (current == 'C') && (next == 'M'):
			result += 900
			i++
		case current == 'M':
			result += 1000
		case (current == 'C') && (next == 'D'):
			result += 400
			i++
		case current == 'D':
			result += 500
		case (current == 'X') && (next == 'C'):
			result += 90
			i++
		case current == 'C':
			result += 100
		case (current == 'X') && (next == 'L'):
			result += 40
			i++
		case current == 'L':
			result += 50
		case (current == 'I') && (next == 'X'):
			result += 9
			i++
		case current == 'X':
			result += 10
		case (current == 'I') && (next == 'V'):
			result += 4
			i++
		case current == 'V':
			result += 5
		case current == 'I':
			result++
		}
	}

	return result
}
