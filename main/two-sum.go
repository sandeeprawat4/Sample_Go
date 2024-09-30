/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
    var m map[int]int
	m = make (map[int]int)
    for _, j:= range nums{
		var k = m[j]
		//fmt.Println("K value", k, "and current arr is ", i, "and curr val is ", j)
		if k!=0 {
			k++;
			m[j]=k;
		}else{
			m[j]=1;
		}
	}
	var first = -1;
	var second = -1;
	for i, j:= range nums{
		var temp = target - j;
		if(m[temp]!=0){
			if(temp == j && m[temp]==1){
				continue
			}
			if(first==-1){
				first=i;
			}else{
				second=i;
				break;
			}
		}
	}
	return []int{first, second}
}
// @lc code=end

