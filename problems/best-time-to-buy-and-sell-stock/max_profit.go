package besttimetobuyandsellstock

func MaxProfit(prices []int) int {
	var (
		minPrice  int
		maxProfit int
	)

	for index, price := range prices {
		if index == 0 {
			minPrice = price
			maxProfit = 0

			continue
		}

		if price < minPrice {
			minPrice = price

			continue
		}

		if profit := price - minPrice; profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}
