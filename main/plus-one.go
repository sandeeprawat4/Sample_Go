/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	n := len(digits)
	for i:=n-1 ; i>=0 ;i--{
		if digits[i]<9{
			digits[i]++
			return digits
		}
		digits[i]=0
	}
	digits = append([]int{1}, digits...)
	return digits    
}
// @lc code=end

