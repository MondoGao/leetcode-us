/*
 * @lc app=leetcode id=977 lang=rust
 *
 * [977] Squares of a Sorted Array
 */

// @lc code=start
impl Solution {
    pub fn sorted_squares(nums: Vec<i32>) -> Vec<i32> {
        let len = nums.len().try_into().unwrap();
        let mut head: usize = 0;
        let mut tail: usize = len - 1;

        let mut out = vec![0; len];

        for v in out.iter_mut().rev() {
            if let (Some(head_v), Some(tail_v)) = (nums.get(head), nums.get(tail)) {
                let next = if head_v.abs() > tail_v.abs() {
                    head += 1;
                    head_v
                } else {
                    tail -= 1;
                    tail_v
                };

                *v = next.pow(2)
            }
        }

        out
    }
}
// @lc code=end
