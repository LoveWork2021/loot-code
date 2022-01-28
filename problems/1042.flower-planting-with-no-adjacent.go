package main

import "fmt"

// https://leetcode-cn.com/problems/flower-planting-with-no-adjacent

func gardenNoAdj(n int, paths [][]int) []int {
	// 缓存所有关联关系
	m := map[int]map[int]bool{}
	for _, path := range paths {
		a := path[0]
		b := path[1]

		if _, ok := m[a]; ok {
			m[a][b] = true
		} else {
			m[a] = map[int]bool{b: true}
		}

		if _, ok := m[b]; ok {
			m[b][a] = true
		} else {
			m[b] = map[int]bool{a: true}
		}
	}

	as := make([]int, n)
	as[0] = 1
	for i := 2; i <= n; i++ {
		if _, ok := m[i]; !ok {
			as[i-1] = 1
			continue
		}

		f := 0b0000
		for k := range m[i] {
			if k > i {
				continue
			}

			f = as[k-1] | f
		}

		if f&0b0001 == 0 {
			as[i-1] = 0b0001
			continue
		} else if f&0b0010 == 0 {
			as[i-1] = 0b0010
			continue
		} else if f&0b0100 == 0 {
			as[i-1] = 0b0100
			continue
		} else if f&0b1000 == 0 {
			as[i-1] = 0b1000
			continue
		}
	}

	r := make([]int, n)
	for i, a := range as {
		if a == 0b0001 {
			r[i] = 1
			continue
		} else if a == 0b0010 {
			r[i] = 2
			continue
		} else if a == 0b0100 {
			r[i] = 3
			continue
		} else if a == 0b1000 {
			r[i] = 4
			continue
		}
	}
	return r
}

func main() {
	paths := [][]int{}
	paths = append(paths, []int{1, 2})
	paths = append(paths, []int{2, 3})
	paths = append(paths, []int{3, 1})

	r := gardenNoAdj(3, paths)
	fmt.Println(r)
}
