package leetcodedemo

// leetcode 第一题
// 解法一
func twoSum(nums []int, target int) []int {
	s := make([]int, 2)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				s = append(s, j)
				return s
			}
		}
	}
	return nil
}

//func twoSum2(nums []int, target int) []int {
//	s := make([]int, 2)
//	var j int = 0
//
//	return nil
//}
