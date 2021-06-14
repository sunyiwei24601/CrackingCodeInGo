package _448_Find_All_Numbers_Disappeared
func findDisappearedNumbers(nums []int) []int {
	result := make([]int, 0)
	for i := range nums{
		j := nums[i]
		for j != i + 1 && j <= len(nums) && nums[j-1] != j {
			if  nums[j - 1] != j {
				nums[i] = nums[j - 1]
				nums[j - 1] = j
			}
			j = nums[i]
		}


	}
	for i, j := range nums {
		if (j - i) != 1 {
			result = append(result, i + 1)
		}
	}

	return result
}