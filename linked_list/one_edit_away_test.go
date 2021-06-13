package linked_list

import (
	"math"
	"testing"
)

func CheckOneEditAway(first, second string) bool {
	if math.Abs(float64(len(first) - len(second))) > 1 {
		return false
	}
	shorter := first
	longer := second
	if len(first) > len(second) {
		shorter = second
		longer = first
	}

	shorter_index := 0
	longer_index := 0
	show_difference := false

	for shorter_index < len(shorter) {
		if shorter[shorter_index] == longer[longer_index] {
			shorter_index += 1
		} else {
			if show_difference {
				return false
			} else {
				show_difference = true
			}
			if len(shorter) == len(longer) {
				shorter_index ++
			}
		}
		longer_index ++
	}
	return true
}

func TestOneEditAway(t *testing.T) {
	println(CheckOneEditAway("pale", "ple"))
	println(CheckOneEditAway("pales", "pale"))
	println(CheckOneEditAway("pale", "bale"))
	println(CheckOneEditAway("pale", "bae"))

}