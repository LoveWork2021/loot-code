package main

import "fmt"

// https://leetcode-cn.com/problems/cheapest-flights-within-k-stops

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	fm := map[int][][]int{}
	for _, f := range flights {
		from := f[0]

		if _, ok := fm[from]; ok {
			fm[from] = append(fm[from], f[1:])
		} else {
			fm[from] = [][]int{f[1:]}
		}
	}

	// 从 src 一个个遍历
	paths := map[int]int{src: 0}

	// 最多经过 k 个中转站，说明最多飞 k+1 次
	for i := 0; i <= k; i++ {
		// 只保留最近一次的价格
		nextPaths := map[int]int{}
		if price, ok := paths[dst]; ok {
			nextPaths[dst] = price
		}

		for to, price := range paths {
			if to == dst {
				continue
			}

			for _, f := range fm[to] {
				nextTo := f[0]
				nextPrice := f[1] + price

				if v, ok := nextPaths[nextTo]; ok {
					if v > nextPrice {
						nextPaths[nextTo] = nextPrice
					}
				} else {
					nextPaths[nextTo] = nextPrice
				}
			}
		}

		paths = nextPaths
	}

	if p, ok := paths[dst]; ok {
		return p
	} else {
		return -1
	}
}

func main() {
	r := findCheapestPrice(
		3,
		[][]int{[]int{0, 1, 2}, []int{1, 2, 1}, []int{2, 0, 10}},
		1,
		2,
		1,
	)

	fmt.Println(r)
}
