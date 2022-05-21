package twosum

//nolint:revive
func TwoSumInt(nums []int, target int) []int {
	memory := make(map[int]int, 0)

	for i, num := range nums {
		if pos, ok := memory[target-num]; ok {
			return []int{pos, i}
		}

		memory[num] = i
	}

	return nil
}
