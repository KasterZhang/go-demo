package leetcode

//DONE
func isLongPressedName925(name string, typed string) bool {
	limit := len(typed) - len(name)
	skip, count := 0, 0
	nameLen := len(name)
	typedLen := len(typed)
	if typedLen < nameLen || nameLen == 0 || typedLen == 0 || name[0] != typed[0] {
		return false
	}
	var pre rune
	for count+skip < typedLen {
		if count >= nameLen {
			if pre != rune(typed[count+skip]) {
				return false
			}
		} else {
			temp := rune(name[count])
			for _, r := range typed[count+skip:] {
				if temp == r {
					break
				}
				skip++
				if pre == 0 {
					pre = r
				} else if r != pre {
					return false
				} else {
					if skip > limit {
						return false
					}
				}
			}
			pre = temp
		}
		count++
	}
	return true
}
