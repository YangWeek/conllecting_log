package leetcodedemo

func differenceOfSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var x int = 0
	var y int = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 10 {
			x += nums[i]
			//y += nums[i] - (len(strconv.Itoa(nums[i])) * 10)

		} else {
			x += nums[i]
			y += nums[i]
		}
	}
	return 0
}
