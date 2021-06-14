package _076_Minimum_Window_Substring

// #string #sliding window

func MinimumWindowSubstring(s, t string) string {
	left := 0
	right := -1
	sFreq := make(map[int32]int, 0)
	tFreq := make(map[int32]int, 0)
	count := 0 // set count to check if the window contain all the character in t, only character in t can be added in count
	tCount := 0
	minW := len(s) + 1
	finalLeft := -1
	finalRight := -1

	for _, j := range t {
		tFreq[j] += 1
		tCount++
	}

	for left < len(s) {
		if count < tCount && right+1 < len(s) { // move right to next one
			c := int32(s[right + 1])
			sFreq[c]++
			if sFreq[c] <= tFreq[c] {
				count++
			}
			right++
		} else {
			if count == tCount && right-left+1 < minW { // find a smaller window
				minW = right - left + 1
				finalLeft = left
				finalRight = right
			}

			c := int32(s[left])  // move left pointer to next
			sFreq[c]--
			if sFreq[c] < tFreq[c] {
				count--
			}
			left++
		}
	}
	return s[finalLeft : finalRight+1]
}
