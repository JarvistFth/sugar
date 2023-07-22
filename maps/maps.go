package maps

import (
	"github.com/JarvistFth/sugar/types"
	"sort"
)

func Filter[K comparable, T any](m1 map[K]T, f func(K, T) bool) (ret map[K]T) {
	for k, v := range m1 {
		if !f(k, v) {
			delete(m1, k)
		}
	}
	return m1
}

func FilterClone[K comparable, T any](m1 map[K]T, f func(K, T) bool) (ret map[K]T) {
	ret = make(map[K]T)
	for k, v := range m1 {
		if f(k, v) {
			ret[k] = v
		}
	}
	return ret
}

func Map[K1, K2 comparable, T, U any](m1 map[K1]T, f func(K1, T) (K2, U)) (ret map[K2]U) {
	ret = make(map[K2]U)
	for k, v := range m1 {
		k2, v2 := f(k, v)
		ret[k2] = v2
	}
	return ret
}

func Keys[K comparable, T any](m1 map[K]T) (ret []K) {
	ret = make([]K, len(m1))
	i := 0
	for k := range m1 {
		ret[i] = k
		i++
	}
	return ret
}

func Values[K comparable, T any](m1 map[K]T) (ret []T) {
	ret = make([]T, len(m1))
	i := 0
	for _, v := range m1 {
		ret[i] = v
		i++
	}
	return ret
}

func OrderedKeys[K types.Sortable, T any](m1 map[K]T) (ret []K) {
	ret = make([]K, len(m1))
	i := 0

	for k := range m1 {
		ret[i] = k
		i++
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i] < ret[j]
	})
	return ret
}

func OrderedValues[K comparable, T types.Sortable](m1 map[K]T) (ret []T) {
	ret = make([]T, len(m1))
	i := 0

	for _, v := range m1 {
		ret[i] = v
		i++
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i] < ret[j]
	})
	return ret
}

func Union[K comparable, T any](origin map[K]T, others ...map[K]T) map[K]T {
	for _, other := range others {
		for k, v := range other {
			origin[k] = v
		}
	}
	return origin
}

func UnionClone[K comparable, T any](origin map[K]T, others ...map[K]T) (ret map[K]T) {
	ret = map[K]T{}
	for _, other := range others {
		for k, v := range other {
			ret[k] = v
		}
	}
	return ret
}
