package twosum

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

//nolint:revive
func TwoSumGeneric[T Number](nums []T, target T) []int {
	memory := make(map[T]int, 0)

	for i, num := range nums {
		if pos, ok := memory[target-num]; ok {
			return []int{pos, i}
		}

		memory[num] = i
	}

	return nil
}
