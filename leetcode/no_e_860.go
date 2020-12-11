package leetcode

//DONE
func lemonadeChange860S1(bills []int) bool {
	box := make(map[int]int)
	for _, cost := range bills {
		change := cost - 5
		if !changeBack(box, change) {
			return false
		}
		if cost != 20 {
			box[cost]++
		}
	}
	return true
}

func changeBack(box map[int]int, change int) bool {
	if change == 15 {
		if count, ok := box[10]; ok {
			if count == 1 {
				delete(box, 10)
			} else {
				box[10] = count - 1
			}
			if count, ok := box[5]; ok {
				if count == 1 {
					delete(box, 5)
				} else {
					box[5] = count - 1
				}
			} else {
				return false
			}
		} else {
			if count, ok := box[5]; ok && count >= 3 {
				if count == 3 {
					delete(box, 5)
				} else {
					box[5] = count - 3
				}
			} else {
				return false
			}
		}
	} else if change == 5 {
		if count, ok := box[5]; ok {
			if count == 1 {
				delete(box, 5)
			} else {
				box[5] = count - 1
			}
		} else {
			return false
		}
	}
	return true
}

//DONE
func lemonadeChange860S2(bills []int) bool {
	box := make([]int, 2)
	for _, bill := range bills {
		if bill == 10 {
			box[0]--
			if box[0] < 0 {
				return false
			}
			box[1]++
		} else if bill == 20 {
			if box[1] > 0 {
				box[1]--
				box[0]--
				if box[0] < 0 {
					return false
				}
			} else {
				box[0] -= 3
				if box[0] < 0 {
					return false
				}
			}
		} else {
			box[0]++
		}

	}
	return true
}
