func change(amount int, coins []int) int {
	denoms := make([]int, amount+1)
	denoms[0] = 1 // 1 way aka not selecting any coins
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			denoms[j] += denoms[j-coins[i]] // add possible way to generate amount using lower denoms
		}
	}
	return denoms[amount]
}