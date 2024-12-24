/*
 * @lc app=leetcode id=977 lang=golang
 *
 * [977] Squares of a Sorted Array
 */
package main

import (
	"fmt"
	"math"
)

// @lc code=start
func sortedSquares(nums []int) []int {
	var (
		in_len     = len(nums)
		out        = make([]int, in_len)
		head, tail = 0, in_len - 1
		cur        = in_len
	)

	for head <= tail {
		cur--
		head_v := float64(nums[head])
		tail_v := float64(nums[tail])
		if math.Abs(float64(head_v)) > math.Abs(float64(tail_v)) {
			out[cur] = int(math.Pow(head_v, 2))
			head++
			continue
		}

		out[cur] = int(math.Pow(tail_v, 2))
		tail--
	}
	return out
}

// @lc code=end
