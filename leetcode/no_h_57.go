package leetcode

// 给出一个无重叠的 ，按照区间起始端点排序的区间列表。

// 在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

//

// 示例 1：

// 输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
// 输出：[[1,5],[6,9]]
// 示例 2：

// 输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
// 输出：[[1,2],[3,10],[12,16]]
// 解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
//

// 注意：输入类型已在 2019 年 4 月 15 日更改。请重置为默认代码定义以获取新的方法签名。

//DONE
//TODO 可以改为O(1)空间 使用合并的概念处理
func insert57(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	if newInterval[1] < intervals[0][0] {
		return append([][]int{newInterval}, intervals...)
	}
	result := [][]int{}
	start, end := 0, 0
	findStart, findEnd := false, false
	for i, interval := range intervals {
		if interval[1] < newInterval[0] {
			// fmt.Println("result",result,"add",interval)
			result = append(result, interval)
		}
		if !findStart {
			if interval[1] >= newInterval[0] {
				findStart = true
				if interval[0] <= newInterval[0] {
					start = interval[0]
				} else {
					start = newInterval[0]
				}
			}
		}

		if findStart && !findEnd {
			if interval[0] > newInterval[1] {
				findEnd = true
				end = newInterval[1]
				result = append(result, []int{start, end})
				// fmt.Println("result",result,"start",start,"end",end)
			} else if interval[1] >= newInterval[1] {
				findEnd = true
				end = interval[1]
				result = append(result, []int{start, end})
				// fmt.Println("result",result,"start",start,"end",end)
			}
		}
		if interval[0] > newInterval[1] {
			// fmt.Println("result",result,"add",interval)
			result = append(result, interval)
		}
		if i == len(intervals)-1 {
			if !findStart {
				// fmt.Println("result",result,"add",newInterval)
				result = append(result, newInterval)
			}
			if findStart && !findEnd {
				findEnd = true
				end = interval[1]
				if newInterval[1] > end {
					end = newInterval[1]
				}
				result = append(result, []int{start, end})
				// fmt.Println("result",result,"start",start,"end",end)
			}
		}
	}
	return result
}
