package slice

import "testing"

func TestCheckIndex(t *testing.T) {
	for i, test := range []struct {
		slice interface{}
		index int
		ret   bool
	}{
		{[]int(nil), 0, false},
		{[]int{}, 0, false},
		{[]int{1, 2}, 0, true},
		{[]int{1, 2}, 1, true},
		{[]int{1, 2}, 2, false},
		{[]int{1, 2}, -1, false},
		{[][]int{}, 0, false},
		{[][]int{{}}, 0, true},
		{[][]int{{}}, 1, false},
	} {
		if ret := CheckIndex(test.slice, test.index); ret != test.ret {
			t.Fatalf("case:%v slice:%#v index:%v require:%v, but:%v",
				i, test.slice, test.index, test.ret, ret)
		}
	}
}
