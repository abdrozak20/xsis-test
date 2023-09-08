package config

import "strconv"

func StringToInt64(buf string, dval int64) int64 {
	res, err := strconv.ParseInt(buf, 10, 64)

	if err == nil {
		return res
	}

	return dval
}

func Int64ToString(val int64) string {
	return strconv.FormatInt(val, 10)
}

func Int16ToString(val int16) string {
	return strconv.FormatInt(
		int64(val),
		10,
	)
}
