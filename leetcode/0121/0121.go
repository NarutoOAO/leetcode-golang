package _121

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	min, result := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > result {
			result = prices[i] - min
		}
		if min > prices[i] {
			min = prices[i]
		}
	}
	return result
}
