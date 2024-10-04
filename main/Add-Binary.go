/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start
func addBinary(a string, b string) string {
	carry := 0
	A := []rune(a)
	B := []rune(b)
	slices.Reverse(A)
	slices.Reverse(B)
	res := []rune{}
	for i:=0 ; i<len(A) || i<len(B) ; i++{
		if i<len(A){
			tmp := int(A[i]-'0')
			carry += tmp;
		}
		if i<len(B){
			tmp := int(B[i]-'0')
			carry += tmp;
		}
		if carry%2==0{
			res = append(res, '0')
		}else{
			res = append(res, '1')
		}
		carry = carry/2
	}
	if carry>0{
		res = append(res, '1')
	}
	slices.Reverse(res)
	r := string(res)
	return r
}
// @lc code=end

