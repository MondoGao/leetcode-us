/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */
 package main

 // @lc code=start
 func reverseString(s []byte) {
   var (
     head int
     tail int = len(s) - 1
   )
 
   for ; head < tail; head, tail = head+1, tail-1 {
     s[tail], s[head] = s[head], s[tail]
   }
 }
 
 // @lc code=end
 