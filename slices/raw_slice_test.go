package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"sugar/choose"
	"sugar/conv"
)

func TestFilter(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := Filter(s1, func(i int) bool {
		if i == 2 {
			return false
		}
		return true
	})
	s3 := []int{1, 3, 4, 5}
	assert.Equal(t, s2, s3)
}

func TestMap(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := Map(s1, func(t int) string {
		return strconv.FormatInt(int64(t), 10)
	})

	s3 := []string{"1", "2", "3", "4", "5"}
	assert.Equal(t, s2, s3)
}

func TestReduce(t *testing.T) {

	s1 := []int{1, 2, 3, 4, 5}

	val := Reduce(s1, func(i int, i2 int) int {
		return choose.Max(i, i2)
	})

	assert.Equal(t, val, 5)
}

func TestRemove(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := Remove(s1, 2, 4)
	assert.Equal(t, s2, []int{1, 3, 5, 6})
}

func TestContains(t *testing.T) {
	assert.True(t, Contains([]int{1, 2, 3, 4, 5}, 5))
}

func TestEqual(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}

	assert.True(t, Equal(s1, s2))

	s3 := []int{1, 2, 3, 4}
	s4 := []int{1, 2, 3, 4, 5}

	assert.False(t, Equal(s3, s4))

}

func TestEqualWith(t *testing.T) {

	type Person struct {
		Name string
		ID   int
	}
	s1 := []Person{
		{
			ID:   1,
			Name: "a",
		},
		{
			ID:   2,
			Name: "b",
		},
	}
	s2 := []Person{
		{
			ID:   1,
			Name: "c",
		},
		{
			ID:   2,
			Name: "d",
		},
	}
	assert.True(t, EqualWith(s1, s2, func(v1, v2 Person) bool {
		return v1.ID == v2.ID || v1.Name == v2.Name
	}))

	assert.False(t, EqualWith(s1, s2, func(v1, v2 Person) bool {
		return v1.ID == v2.ID && v1.Name == v2.Name
	}))
}

func TestToMap(t *testing.T) {
}

func TestToMapValue(t *testing.T) {

	s1 := []int{1, 2, 3, 4, 5}
	m1 := ToMapValue(s1, func(t int) (k string) {
		return fmt.Sprintf("key:%d", t)
	})

	for k, v := range m1 {
		assert.Equal(t, fmt.Sprintf("key:%d", v), k)
	}
}

func TestForEach(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6}
	i := 0
	ForEach(s1, func(d int) {
		assert.Equal(t, i+1, d)
		i++
	})
}

func TestChunk(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6}

	s2 := Chunk(s1, 2)
	t.Logf("s2=%+v", s2)
	assert.Equal(t, len(s2), 3)
	assert.Equal(t, s2[0][0], 1)
	assert.Equal(t, s2[0][1], 2)
	assert.Equal(t, s2[1][0], 3)
	assert.Equal(t, s2[1][1], 4)
	assert.Equal(t, s2[2][0], 5)
	assert.Equal(t, s2[2][1], 6)

	s1 = []int{1, 2, 3, 4, 5}

	s2 = Chunk(s1, 3)
	t.Logf("s2=%+v", s2)
	assert.Equal(t, len(s2), 2)
	assert.Equal(t, s2[0][0], 1)
	assert.Equal(t, s2[0][1], 2)
	assert.Equal(t, s2[0][2], 3)
	assert.Equal(t, s2[1][0], 4)
	assert.Equal(t, s2[1][1], 5)

}

func TestFlatMap(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := FlatMap[int, string](s1, func(i int) []string {
		return []string{conv.IntToString(i), conv.IntToString(i + 10)}
	})
	s3 := []string{"1", "11", "2", "12", "3", "13", "4", "14", "5", "15"}
	assert.True(t, Equal(s2, s3))
}

func TestFlatten(t *testing.T) {
	s1 := [][]int{{1, 2, 3, 4, 5}, {5, 4, 3, 2, 1}, {100, 200}, {400, 200}}
	s2 := Flatten(s1)
	s3 := []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1, 100, 200, 400, 200}
	assert.True(t, Equal(s2, s3))
}

func TestAnyOf(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	assert.True(t, AnyOf(s1, func(i int) bool {
		return i == 2
	}))
	assert.False(t, AnyOf(s1, func(i int) bool {
		return i > 6
	}))
}

func TestAllOf(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	assert.True(t, AllOf(s1, func(i int) bool {
		return i < 6
	}))
	assert.False(t, AllOf(s1, func(i int) bool {
		return i%2 == 0
	}))
}

func TestReverse(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := Reverse(s1)
	s3 := []int{6, 5, 4, 3, 2, 1}
	assert.True(t, Equal(s1, s2))
	assert.True(t, Equal(s2, s3))
}

func TestReverseClone(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := ReverseClone(s1)
	s3 := []int{6, 5, 4, 3, 2, 1}
	assert.False(t, Equal(s1, s2))
	assert.True(t, Equal(s2, s3))
}
