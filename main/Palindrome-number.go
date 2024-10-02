/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x<0{
		return false
	}
    num := x
	revNum :=0

	for x>0 {
		revNum = revNum * 10 + x%10
		x = x/10;
	}
	return num == revNum
}
// @lc code=end