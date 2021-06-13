package linked_list

import (
	"fmt"
	"testing"
)

func IsUnique(s string) bool {
	if len(s) > 128 {
		return false
	}

	var result [128]bool
	for i := 0; i < len(s); i++ {
		c := int(s[i])
		if result[c] {
			return false
		} else {
			result[c] = true
		}
	}
	return true
}

func TestIsUnique(t *testing.T) {
	fmt.Println(IsUnique("abcda"))
	fmt.Println(IsUnique("abcd"))
}
