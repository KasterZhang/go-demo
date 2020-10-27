package leetcode

func videoStitching1024(clips [][]int, T int) int {
	step, start, end := 0, 0, T
	if T == 0 {
		return 0
	}
	if len(clips) == 0 {
		return -1
	}
	for start < end {
		// fmt.Printf("--------- start %d end %d ----------",start, end)
		ts, te, k, j := 0, 0, -1, -1
		for i, clip := range clips {
			if clip[0] <= start && clip[1] > start {
				if clip[1] >= end {
					return step + 1
				}
				// fmt.Println("start", start, "ts", ts, "clip", clip)
				if clip[1]-start > ts {
					ts = clip[1] - start
					k = i
				}
			}
			if clip[1] >= end && clip[0] < end {
				// fmt.Println("end", end, "te", te, "clip", clip)
				if end-clip[0] > te {
					te = end - clip[0]
					j = i
				}
			}
		}
		if ts == 0 || te == 0 {
			return -1
		}
		start += ts
		end -= te
		step++
		if k != j {
			step++
		}
	}
	return step
}
