package leetcode

//OrderedStream1656 DONE
type OrderedStream1656 struct {
	dict []string
	ptr  int
}

//Constructor1656 DONE
func Constructor1656(n int) OrderedStream1656 {
	dict := make([]string, n)
	return OrderedStream1656{
		dict: dict,
		ptr:  1,
	}
}

//Insert DONE
func (o *OrderedStream1656) Insert(id int, value string) []string {
	o.dict[id-1] = value
	// fmt.Println("prt",o.ptr,"id",id,"dict",o.dict)
	if id == o.ptr {
		for _, n := range o.dict[id-1:] {
			// fmt.Println("ptr move",i,n,o.ptr)
			if n == "" {
				return o.dict[id-1 : o.ptr-1]
			}
			o.ptr++
		}
		return o.dict[id-1:]
	}
	return []string{}
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(id,value);
 */
