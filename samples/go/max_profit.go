package main

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func maxProfit(prices []int) int {
	localMax := make([]int, len(prices))
	maxPrice := 0
	for i := len(prices) - 1; i > -1; i-- {
		maxPrice = max(maxPrice, prices[i])
		localMax[i] = maxPrice
	}

	maxDiff := 0
	for i := range prices {
		if localMax[i]-prices[i] > maxDiff {
			maxDiff = localMax[i] - prices[i]
		}
	}
	return maxDiff
}
