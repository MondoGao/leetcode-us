#
# @lc app=leetcode id=977 lang=python3
#
# [977] Squares of a Sorted Array
#

# @lc code=start
from typing import List


class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        out = [0] * len(nums)
        head = 0
        tail = len(nums) - 1
        cur = tail

        while head <= tail:
            next = 0
            if abs(nums[head]) > abs(nums[tail]):
                next = nums[head]
                head += 1
            else:
                next = nums[tail]
                tail -= 1

            out[cur] = next**2
            cur -= 1
        
        return out


# @lc code=end
