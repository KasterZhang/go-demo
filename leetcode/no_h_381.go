package leetcode

import (
	"math/rand"
)

// 设计一个支持在平均 时间复杂度 O(1) 下， 执行以下操作的数据结构。

// 注意: 允许出现重复元素。

//RandomizedCollection leetcode
// insert(val)：向集合中插入元素 val。
// remove(val)：当 val 存在时，从集合中移除一个 val。
// getRandom：从现有集合中随机获取一个元素。每个元素被返回的概率应该与其在集合中的数量呈线性相关。
//DONE
type RandomizedCollection381 struct {
	arr  []int
	dict map[int]map[int]byte
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection381 {
	return RandomizedCollection381{
		arr:  []int{},
		dict: make(map[int]map[int]byte),
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection381) Insert(val int) bool {
	this.arr = append(this.arr, val)
	if _, ok := this.dict[val]; ok {
		this.dict[val][len(this.arr)-1] = byte(1)
		return false
	}
	this.dict[val] = map[int]byte{len(this.arr) - 1: byte(1)}
	return true
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection381) Remove(val int) bool {
	count := len(this.arr)
	if indexs, ok := this.dict[val]; ok {
		last := this.arr[count-1]
		if val == last {
			this.arr = this.arr[:count-1]
			delete(this.dict[last], len(this.arr))
		} else {
			index := -1
			for k := range indexs {
				index = k
				break
			}
			//arr
			this.arr[index] = last
			this.arr = this.arr[:len(this.arr)-1]
			//map
			this.dict[last][index] = byte(1)
			delete(this.dict[last], len(this.arr))
			delete(this.dict[val], index)

		}

		if len(this.dict[val]) == 0 {
			delete(this.dict, val)
		}
		return true
	}
	return false
}

/** Get a random element from the collection. */
func (this *RandomizedCollection381) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}
