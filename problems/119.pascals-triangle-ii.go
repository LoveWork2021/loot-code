package main

import "fmt"

// https://leetcode-cn.com/problems/pascals-triangle-ii

func getRow(rowIndex int) []int {
	l := make([]int, rowIndex+1)
	l[0] = 1

	for r := 1; r <= rowIndex; r++ {
		for c := r; c > 0; c-- {
			l[c] += l[c-1]
		}
	}

	return l
}

func main() {
	r := getRow(3)
	fmt.Println(r)
}
