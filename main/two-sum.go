func twoSum(nums []int, target int) []int {

	m := make(map[int]int)
    for index, arrVal:= range nums{
		diff := target - arrVal
		if sIndex, found := m[diff]; found{
			return []int{index, sIndex}
		}
		m[arrVal] = index;
	}
	return nil
}