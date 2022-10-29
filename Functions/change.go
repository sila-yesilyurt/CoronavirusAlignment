package Functions

//Change takes an amount of money along with a collection of denominations.
//It returns the minimum number of coins needed to change the money given the denominations.
func Change(money int, Coins []int) int {
	minNumCoins := make([]int, money+1)
	for k := 1; k <= money; k++ {
		// take minimum of all relevant values
		var currentMin int
		for i := range Coins {
			if k-Coins[i] >= 0 {
				if i == 0 || minNumCoins[k-Coins[i]] < currentMin {
					currentMin = minNumCoins[k-Coins[i]]
				}
			}
		}
		minNumCoins[k] = currentMin + 1
	}
	return minNumCoins[money]
}
