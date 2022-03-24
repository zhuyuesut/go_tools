package slice

import "reflect"

func CheckIndex(slice interface{}, i int) bool {
	return i >= 0 && i < reflect.ValueOf(slice).Len()
}
