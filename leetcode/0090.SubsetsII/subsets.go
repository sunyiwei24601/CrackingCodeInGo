package _090_SubsetsII

import "sort"

func subsetsWithDup(nums []int)[][]int {
	sort.Ints(nums)
	count := make(map[int]int, 0)
	for _, num := range nums {
		count[num] ++
	}
	result := make([][]int, 0)
	result = append(result, []int{})

	for key, num := range count {
		currentResult := addK(result, num, key) // 计算出增加当前元素后的结果
		for _, current := range currentResult {
			result = append(result, current)
		}
	}
	return result
}

func addK(result [][]int, k int, a int) [][]int{
	// 给出一个数字a以及它出现的次数k，在原有的结果中，往里加1个这个数字，
	// 两个...直到k个，统计所有的情况在currentResult里， 简单来说就是将同一个数字出现多次视为一个单独的数字
	currentResult := make([][]int, 0)
	for times:=1; times <= k; times++{
		for _, current := range result {
			temp := make([]int, len(current))
			copy(temp, current)
			for i := 0; i < times; i++ {
				temp = append(temp, a)
			}
			currentResult = append(currentResult, temp)
		}
	}
	return currentResult
}

func subsetsWithDup2(nums []int)[][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for maxLength:=0; maxLength<=len(nums); maxLength++{
		generateSubsets(nums,  maxLength, 0, []int{}, &result)
	}
	return result
}

func generateSubsets(nums []int, maxLength int, start int, currentSet []int, result *[][]int){
	if len(currentSet) == maxLength {
		c := make([]int, len(currentSet))
		copy(c, currentSet)
		*result = append(*result, c)
		return
	}

	for i:=start; i < len(nums) - (maxLength - len(currentSet)) + 1; i++ {
		if i > start && nums[i] == nums[i-1] { //去重
			continue
		}
		currentSet = append(currentSet, nums[i])
		generateSubsets(nums, maxLength, i + 1, currentSet, result)
		currentSet = currentSet[:len(currentSet) - 1]
	}
}