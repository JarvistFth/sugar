package slices

func Filter[T any](arr []T, f func(T) bool) (ret []T) {
	for _, v := range arr {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func Map[T, U any](arr []T, f func(T) U) (ret []U) {
	for _, v := range arr {
		ret = append(ret, f(v))
	}
	return ret
}

func Reduce[T any](arr []T, f func(T, T) T) (res T) {
	if len(arr) <= 0 {
		return
	}
	prev := arr[0]
	for _, v := range arr {
		res = f(prev, v)
		prev = v
	}
	return res
}

func Remove[T comparable](arr []T, i ...T) (ret []T) {
	return Filter(arr, func(t T) bool {
		for _, ii := range i {
			if ii == t {
				return false
			}
		}
		return true
	})
}

func Contains[T comparable](arr []T, i T) bool {
	for _, v := range arr {
		if v == i {
			return true
		}
	}

	return false
}

func Equal[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func EqualWith[T comparable](s1, s2 []T, f func(v1, v2 T) bool) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if !f(s1[i], s2[i]) {
			return false
		}
	}

	return true
}

func ToMap[K comparable, V, T any](arr []T, f func(t T) (k K, v V)) (ret map[K]V) {
	ret = make(map[K]V)

	for _, vv := range arr {
		k, v := f(vv)
		ret[k] = v
	}

	return ret
}

func ToMapValue[K comparable, T any](arr []T, f func(t T) (k K)) (ret map[K]T) {
	ret = make(map[K]T)

	for _, vv := range arr {
		k := f(vv)
		ret[k] = vv
	}

	return ret
}

func ForEach[T any](arr []T, f func(t T)) {
	for _, vv := range arr {
		f(vv)
	}
}

func Chunk[T any](arr []T, n int) (ret [][]T) {
	for i := 0; i < len(arr); {
		if i+n > len(arr) {
			ret = append(ret, arr[i:])
		} else {
			ret = append(ret, arr[i:i+n])
		}
		i = i + n
	}
	return ret
}

func FlatMap[T, U any](arr []T, f func(T) []U) (ret []U) {
	for _, vv := range arr {
		ret = append(ret, f(vv)...)
	}
	return ret
}

func Flatten[T any](arr [][]T) (ret []T) {
	for _, v := range arr {
		ret = append(ret, v...)
	}
	return ret
}

func AnyOf[T any](arr []T, f func(T) bool) (ret bool) {
	for _, vv := range arr {
		if f(vv) {
			return true
		}
	}
	return false
}

func AllOf[T any](arr []T, f func(T) bool) (ret bool) {
	for _, vv := range arr {
		if !f(vv) {
			return false
		}
	}

	return true
}

func Reverse[T any](arr []T) []T {
	for from, to := 0, len(arr)-1; from < to; from, to = from+1, to-1 {
		arr[from], arr[to] = arr[to], arr[from]
	}
	return arr
}

func ReverseClone[T any](arr []T) (ret []T) {
	for i := len(arr) - 1; i >= 0; i-- {
		ret = append(ret, arr[i])
	}
	return ret
}
