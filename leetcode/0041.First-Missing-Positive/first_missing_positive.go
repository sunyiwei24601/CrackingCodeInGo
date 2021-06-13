package _041_First_Missing_Positive

//  #array

func firstMissingPositive(nums []int) int {
	for i := range nums {
		for nums[i] != i+1 && nums[i] >= 0 && nums[i] <= len(nums) {
			j := nums[i]
			if nums[i] == nums[j-1] {
				break
			}
			nums[i] = nums[j-1]
			nums[j-1] = j
		}
	}
	for i, j := range nums {
		if i != j-1 {
			return i + 1
		}
	}
	return -1
}
