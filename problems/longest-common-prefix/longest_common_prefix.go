package longestcommonprefix

func LongestCommonPrefix(strs []string) string {
	var minLen int

	for i, str := range strs {
		l := len(str)
		if i == 0 {
			minLen = l
		} else if l < minLen {
			minLen = l
		}
	}

	for pos := 0; pos < minLen; pos++ {
		var char uint8

		for i, str := range strs {
			if i == 0 {
				char = str[pos]
			} else if char != str[pos] {
				return strs[0][:pos]
			}
		}
	}

	return strs[0][:minLen]
}
