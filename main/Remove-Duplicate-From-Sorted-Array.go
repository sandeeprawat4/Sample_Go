/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
    var ans int
	ans = 1
	for i:=1 ; i<len(nums) ; i++{
		if nums[i]!=nums[i-1]{
			nums[ans]=nums[i]
			ans++
		}
	}
	return ans
}
// @lc code=end

