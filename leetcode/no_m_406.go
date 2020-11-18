package leetcode

//DONE
func reconstructQueue406(people [][]int) [][]int {
	result := [][]int{}
	tempPeople := make([][]int, len(people))
	for i, person := range people {
		tempPerson := make([]int, 2)
		copy(tempPerson, person)
		tempPeople[i] = tempPerson
	}

	for len(people) != 0 {
		index := -1
		for i, person := range tempPeople {
			if person[1] == 0 {
				if index != -1 && tempPeople[index][0] > person[0] {
					index = i
				} else if index == -1 {
					index = i
				}
			}
		}

		temp := people[index]
		result = append(result, temp)
		tempPeople = append(tempPeople[:index], tempPeople[index+1:]...)
		people = append(people[:index], people[index+1:]...)
		for _, person := range tempPeople {
			if person[0] <= temp[0] {
				person[1]--
			}
		}
	}
	return result
}
