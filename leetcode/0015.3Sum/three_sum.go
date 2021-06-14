package _015_3Sum

import "sort"

// #array
func ThreeSum(s []int)[][]int {
	result := make([][]int, 0)
	uniqNums := make([]int, 0)
	count := make(map[int]int, 0)
	for _, i := range s{
		count[i]++
	}
	for key := range count {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)
	for i:=0; i < len(uniqNums); i++ {
		if uniqNums[i] * 3 == 0 && count[uniqNums[i]] >= 3 { // use left 3 times
			result = append(result, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++{
			if (uniqNums[i] * 2 + uniqNums[j]) == 0 && count[uniqNums[i]] >= 2{ //use 2 times left
				result = append(result, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			} else if (uniqNums[j] * 2 + uniqNums[i]) == 0 && count[uniqNums[j]] >= 2{ //use 2 times right
				result = append(result, []int{uniqNums[j], uniqNums[j], uniqNums[i]})
			} else {
				k := 0 - uniqNums[i] - uniqNums[j]
				if k > uniqNums[j] && count[k] > 0 { // don't use any number twice
					result = append(result, []int{uniqNums[i], uniqNums[j], k})
				}



			}




		}



	}




	return result
}
