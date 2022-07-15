package utils

import "ds/datatype"

func Interface2Val(i interface{}) interface{} {
	var val interface{}
	switch v := i.(type) {
	case int:
		val = v
	case *datatype.TNode:
		val = v.Key
	default:
		val = ""
	}
	return val
}
