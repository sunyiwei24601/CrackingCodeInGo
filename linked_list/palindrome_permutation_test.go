package linked_list

import (
	"fmt"
	"testing"
)

func IsPermutationOfPalindrome(s string) bool {
	a := 0
	for _, i := range s {
		c := int(i) - int("a"[0])
		c = 1 << uint(c)
		if a & c == 0 {  // a & 000010000 shows if the same digit position of a is 1
			a |= c
		} else {
			a &= ^c
		}
	}

	if a == 0 || a & (a - 1) == 0{ // a & (a -1) shows there is only one digit set as 1
		return true
	}

	return false
}


func TestIsPermutationOfPalindrome(t *testing.T) {
	fmt.Println(IsPermutationOfPalindrome("a"))
	fmt.Println(IsPermutationOfPalindrome("aba"))
	fmt.Println(IsPermutationOfPalindrome("abc"))

}