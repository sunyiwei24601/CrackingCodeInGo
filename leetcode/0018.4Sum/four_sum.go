package _018_4Sum

import "sort"

// #array
func fourSum(s []int,  target int) [][]int{
	result := make([][]int, 0)
	sort.Ints(s)
	for i:=0; i < len(s) - 3; i++{
		if i > 0 && s[i] == s[i-1] {
			continue
		}
		I := s[i]
		for j := i + 1; j < len(s) - 2; j++{
			if j > i + 1 && s[j] == s[j-1] {
				continue
			}
			J := s[j]
			for left, right := j+1, len(s) - 1; left < right; {
				if s[left] + s[right] + I + J < target {
					left ++
				} else if s[left] + s[right] + I + J > target {
					right --
				} else {
					result = append(result, []int{I, J, s[left], s[right]})
					for left < len(s) - 1 && s[left] == s[left + 1]{
						left ++
					}
					for right > 0 && s[right] == s[right - 1]{
						right --
					}
					left++
				}
			}
		}
	}
	return result
}


