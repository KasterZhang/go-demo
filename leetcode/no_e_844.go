package leetcode

import (
	"fmt"
)

//DONE O(n) o(n)
func backspaceCompare844S1(S string, T string) bool {
	cs := 0
	rs := ""
	for i := len(S) - 1; i >= 0; i-- {
		// fmt.Println("i",i,"btye",string(S[i]))
		if S[i] != byte('#') && cs == 0 {
			rs += string(S[i])
		} else if S[i] != byte('#') && cs > 0 {
			cs--
		} else {
			cs++
		}
	}
	// fmt.Println("rs",rs)
	ct := 0
	rt := ""
	for i := len(T) - 1; i >= 0; i-- {
		// fmt.Println("j",i,"btye",string(T[i]))
		if T[i] != byte('#') && ct == 0 {
			rt += string(T[i])
		} else if T[i] != byte('#') && ct > 0 {
			ct--
		} else {
			ct++
		}
	}
	// fmt.Println("rt",rt)
	return rs == rt
}

//DONE O(n) o(1) 双指针
func backspaceCompare844S2(S string, T string) bool {
	i := len(S) - 1
	j := len(T) - 1
	cs, ct := 0, 0
	for {
		for i >= 0 {
			if S[i] == '#' {
				cs++
			} else if cs != 0 {
				cs--
			} else {
				break
			}
			i--
		}
		fmt.Println(i)
		for j >= 0 {
			if T[j] == '#' {
				ct++
			} else if ct != 0 {
				ct--
			} else {
				break
			}
			j--
		}
		fmt.Println(j)
		if i >= 0 && j >= 0 {
			if S[i] != T[j] {
				return false
			}
		} else if i >= 0 {
			if S[i] != '#' {
				return false
			}
		} else if j >= 0 {
			if T[j] != '#' {
				return false
			}
		} else {
			return true
		}

		i--
		j--
		fmt.Println("-----------")
	}
}
