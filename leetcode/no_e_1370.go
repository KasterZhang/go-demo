package leetcode

import (
	"sort"
	"strings"
)

type runeList []rune

func (r runeList) Len() int { return len(r) }
func (r runeList) Less(i, j int) bool {
	if r[i] < r[j] {
		return true
	}
	return false
}
func (r runeList) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

//DONE
func sortString1370S1(s string) string {
	ss := runeList(s)
	sort.Sort(ss)

	dict := make(map[rune]int)
	level := 0
	d := []rune{}
	for _, k := range ss {
		if count, ok := dict[k]; ok {
			dict[k] = count + 1
		} else {
			d = append(d, k)
			dict[k] = 1
		}
		if dict[k] > level {
			level = dict[k]
		}
	}

	var sb strings.Builder
	for i := 0; i < level; i++ {
		if i&1 == 0 {
			j := 0
			for j < len(d) {
				// fmt.Println("正序",string(d[j]),"dict",dict)
				sb.WriteRune(d[j])
				if dict[d[j]] == 1 {
					delete(dict, d[j])
					d = append(d[:j], d[j+1:]...)
				} else {
					dict[d[j]]--
					j++
				}
			}
		} else {
			for j := len(d) - 1; j >= 0; j-- {
				// fmt.Println("倒序",string(d[j]),"dict",dict)
				sb.WriteRune(d[j])
				if dict[d[j]] == 1 {
					delete(dict, d[j])
					d = append(d[:j], d[j+1:]...)
				} else {
					dict[d[j]]--
				}
			}
		}
		// fmt.Println("----------")
	}
	return sb.String()
}

//DONE
func sortStringS2(s string) string {
	letter := make([]int, 26)
	for _, n := range s {
		letter[n-'a']++
	}
	var sb strings.Builder
	i := 0
	for i < len(s) {
		for j := 0; j < 26; j++ {
			if letter[j] != 0 {
				sb.WriteRune(rune('a' + j))
				letter[j]--
				i++
			}
		}
		for j := 25; j >= 0; j-- {
			if letter[j] != 0 {
				sb.WriteRune(rune('a' + j))
				letter[j]--
				i++
			}
		}
	}

	return sb.String()
}
