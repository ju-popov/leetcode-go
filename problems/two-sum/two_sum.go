package twosum

func TwoSum(nums []int, target int) []int {
	memory := make(map[int]int) // num:pos

	for i, num := range nums {
		if pos, ok := memory[target-num]; ok {
			return []int{pos, i}
		}

		memory[num] = i
	}

	return nil
}
