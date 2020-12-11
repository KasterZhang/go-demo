package leetcode

import (
	"strings"
)

//DONE
func predictPartyVictory649(senate string) string {
	var winner rune
	vote := 0

	var sb strings.Builder

	for {
		for _, s := range senate {
			if s == winner {
				sb.WriteRune(s)
				vote++
			} else {
				if vote == 0 {
					winner = s
					vote++
					sb.WriteRune(s)
				} else if vote > 0 {
					vote--
				} else {
					sb.WriteRune(s)
				}
			}
		}
		if senate == sb.String() {
			break
		}
		senate = sb.String()
	}

	if winner == 'R' {
		return "Radiant"
	}
	return "Dire"
}
