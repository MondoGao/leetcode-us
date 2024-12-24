/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */
 package main

 // @lc code=start
 func reverseString(s []byte) {
   var (
     len  = len(s)
     stop = len / 2
     head int
     tail int = len - 1
   )
 
   for ; head < stop; head, tail = head+1, tail-1 {
     temp := s[tail]
     s[tail] = s[head]
     s[head] = temp
   }
 }
 
 // @lc code=end
 