package main

import "fmt"

func coinChange(total int, coins []int) int {
	dp := make([]int, total+1) // Create an array of size total+1 with initial value 0
	dp[0] = 1                  // Initialize the first element as 1, since there's one way to make 0 money

	for _, coin := range coins {
		for i := coin; i <= total; i++ {
			dp[i] += dp[i-coin] // Accumulate ways of making current amount using current coin
		}
	}
	return dp[total] // Return the total number of ways to make the given total amount
}

func main() {

	total := 10
	coins := []int{1, 5}
	output := coinChange(total, coins)
	fmt.Println(output)

	total2 := 4
	coins2 := []int{1, 2, 3}
	output2 := coinChange(total2, coins2)
	fmt.Println(output2)

}
