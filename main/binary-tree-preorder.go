/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    var nums = []int{}
	if root == nil{
		return nums
	}
	nums = append(nums, root.Val)
	nums = append(nums, preorderTraversal(root.Left)...)
	nums = append(nums, preorderTraversal(root.Right)...)
	return nums
}
// @lc code=end

