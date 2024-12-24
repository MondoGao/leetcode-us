/*
 * @lc app=leetcode id=344 lang=rust
 *
 * [344] Reverse String
 */
// @lc code=start
impl Solution {
    pub fn reverse_string(s: &mut Vec<char>) {
        let mut head: usize = 0;
        let mut tail: usize = s.len() - 1;

        while head < tail {
            s.swap(head, tail);
            head += 1;
            tail -= 1;
        }
    }
}
// @lc code=end

