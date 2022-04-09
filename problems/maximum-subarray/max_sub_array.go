package maximumsubarray

func MaxSubArray(nums []int) int {
	var (
		sum    int
		maxSum int
	)

	for index, num := range nums {
		if index == 0 {
			sum = num
			maxSum = num

			continue
		}

		if sum < 0 {
			sum = 0
		}

		sum += num

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}
