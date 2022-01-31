package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/4sum

func fourSum(nums []int, target int) [][]int {
	n := len(nums)

	result := [][]int{}

	if n < 4 {
		return result
	}

	sort.Ints(nums)

	for ai := 0; ai < n-3 && nums[ai]+nums[ai+1]+nums[ai+2]+nums[ai+3] <= target; ai++ {
		if (ai > 0 && nums[ai] == nums[ai-1]) || nums[ai]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}

		for bi := ai + 1; bi < n-2 && nums[ai]+nums[bi]+nums[bi+1]+nums[bi+2] <= target; bi++ {
			if (bi > ai+1 && nums[bi-1] == nums[bi]) || nums[ai]+nums[bi]+nums[n-2]+nums[n-1] < target {
				continue
			}

			l := bi + 1
			r := n - 1

			for l < r {
				t := nums[ai] + nums[bi] + nums[l] + nums[r]

				if t == target {
					result = append(result, []int{nums[ai], nums[bi], nums[l], nums[r]})

					for l++; l < r && nums[l] == nums[l-1]; l++ {

					}
					for r--; l < r && nums[r] == nums[r+1]; r-- {

					}
				}

				if t < target {
					l++
					continue
				}

				if t > target {
					r--
					continue
				}
			}
		}
	}

	return result
}

func main() {
	r := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	fmt.Println(r)
}
