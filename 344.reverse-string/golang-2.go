/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */
package reversestring2

// @lc code=start
func reverseString(s []byte) {
	var (
		head int
		tail int = len(s) - 1
	)

	for ; head < tail; head, tail = head+1, tail-1 {
		temp := s[tail]
		s[tail] = s[head]
		s[head] = temp
	}
}

// @lc code=end
