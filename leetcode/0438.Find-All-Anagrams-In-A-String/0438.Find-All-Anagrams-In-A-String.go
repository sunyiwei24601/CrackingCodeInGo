package _438_Find_All_Anagrams_In_A_String
func findAnagrams(s string, p string) []int {
	if len(s) < len(p){
		return []int{}
	}

	result := make([]int, 0)
	countP := make(map[int32]int, 0)
	countW := make(map[int32]int, 0)
	for _, j := range p {
		countP[j] += 1
	}

	left := 0
	count := 0
	for i, j := range s {
		if countP[j] > 0 {
			left = i
			break
		}
	}
	right := left - 1
	for left <= len(s)-len(p) && right < len(s)-1 {
		if countP[int32(s[left])] == 0 {
			left++
		} else if right-left < len(p)-1 {
			countW[int32(s[right+1])] ++
			if countW[int32(s[right+1])] <= countP[int32(s[right+1])] {
				count++
			}
			right++
		} else {
			countW[int32(s[left])] --
			if countP[int32(s[left])] > 0 && countW[int32(s[left])] < countP[int32(s[left])] {
				count--
			}
			left++
		}
		if count == len(p) && right-left == len(p)-1 {
			result = append(result, left)
		}
	}
	return result
}






