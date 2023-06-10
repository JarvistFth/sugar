package choose

import "sugar/types"

// If use it carefully
// if a,b is an empty pointer's fields, use this function may cause panic
// e.g.
// var (
//
//	ptr = nil
//	i = 0
//
// )
// If(i !=0, i, ptr.element)
// will cause panic
func If[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

func Max[T types.Sortable](a, b T) T {
	return If(a > b, a, b)
}

func Min[T types.Sortable](a, b T) T {
	return If(a < b, a, b)
}

func MinMax[T types.Sortable](a, b T) (min, max T) {
	min = Min(a, b)
	max = Max(a, b)
	return
}
