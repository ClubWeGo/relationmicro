package util

/*
*
[]string -> []interface{}
*/
func StrArrToInterfaceArr(arr []string) []interface{} {
	is := make([]interface{}, len(arr))
	for i, a := range arr {
		is[i] = a
	}
	return is
}
