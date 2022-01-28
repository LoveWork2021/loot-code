package main

import "fmt"

// https://leetcode-cn.com/problems/coin-change

func _coinChange(coins []int, amount int, cache map[int]int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if n, ok := cache[amount]; ok {
		return n
	}

	min := 9223372036854775807
	for _, coin := range coins {
		n := _coinChange(coins, amount-coin, cache)
		if n >= 0 && n < min {
			min = n
		}
	}

	if min == 9223372036854775807 {
		cache[amount] = -1
	} else {
		cache[amount] = min + 1
	}

	return cache[amount]
}

func coinChange(coins []int, amount int) int {
	return _coinChange(coins, amount, map[int]int{})
}

func main() {
	r := coinChange([]int{2}, 3)
	fmt.Println(r)
}

// COMMENT: 动态规划，多个数相加成一个数，最小组合
