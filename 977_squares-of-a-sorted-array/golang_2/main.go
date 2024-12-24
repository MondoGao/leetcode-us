/*
 * @lc app=leetcode id=977 lang=golang
 *
 * [977] Squares of a Sorted Array
 */
package main

// @lc code=start
func abs(num int) int {
	if num > 0 {
		return num
	}

	return -num
}
func sortedSquares(nums []int) []int {
	var (
		in_len     = len(nums)
		out        = make([]int, in_len)
		head, tail = 0, in_len - 1
		cur        = in_len
	)

	for head <= tail {
		cur--
		head_v := nums[head]
		tail_v := nums[tail]
		if abs(head_v) > abs(tail_v) {
			out[cur] = head_v * head_v
			head++
			continue
		}

		out[cur] = tail_v * tail_v
		tail--
	}
	return out
}

// @lc code=end
