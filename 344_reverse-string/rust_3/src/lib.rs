/*
 * @lc app=leetcode id=344 lang=rust
 *
 * [344] Reverse String
 */
// @lc code=start
impl Solution {
    pub fn reverse_string(s: &mut Vec<char>) {
        let len = s.len();
        (0..s.len() / 2).into_iter().for_each(|head| {
            let tail = len - 1 - head;
            s.swap(head, tail)
        });
    }
}
// @lc code=end

