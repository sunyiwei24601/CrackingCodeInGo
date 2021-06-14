package _016_3SumClosest

import (
	"math"
	"sort"
)

// #array
func threeSumClosest(s []int, target int) []int {
	result := make([]int, 0)
	diff := math.MaxInt32
	sort.Ints(s)
	for i, K := range s {
		for left, right := i + 1, len(s) - 1; left < right; {
			sum := s[left] + s[right] + K
			if abs(sum - target) < diff {
				result = []int{s[left], s[right], K}
				diff = abs(sum - target)
			}
			if sum < target {
				left ++
			} else if sum > target {
				right --
			} else{
				return result
			}
		}
	}
	return result
}


func abs(a int) int {
	return int(math.Abs(float64(a)))
}