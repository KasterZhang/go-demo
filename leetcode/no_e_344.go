package leetcode

//DONE
func reverseString344(s []byte) {
	start, end := 0, len(s)-1
	for start != end && start-1 != end {
		temp := s[start]
		s[start] = s[end]
		s[end] = temp
		start++
		end--
	}
	return
}
