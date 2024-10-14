/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
    runes := []rune(s)
    var ans int = 0
    for i:=len(runes)-1; i>=0 ; i--{
        if runes[i]==' '{
            if ans>0{
                return ans
            }else{
                continue
            }
        }else{
            ans++
        }
    }
    return ans
}
// @lc code=end