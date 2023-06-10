package conv

import (
	"strconv"

	"sugar/types"
)

func StringInt64Default(input string, defaultValue int64) int64 {
	ret, err := strconv.ParseInt(input, 10, 64)
	if err == nil {
		return ret
	}
	return defaultValue
}

func StringToInt64E(input string) (int64, error) {
	return strconv.ParseInt(input, 10, 64)
}

func IntToString[T types.Integer](input T) string {
	return strconv.FormatInt(int64(input), 10)
}

func FloatToString[T types.Float](input T) string {
	return strconv.FormatFloat(float64(input), 'E', -1, 64)
}
