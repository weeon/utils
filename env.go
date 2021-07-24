package utils

import (
	"os"
	"strconv"
)

func GetEvn(key string, df ...string) string {
	v := os.Getenv(key)
	if len(v) == 0 && len(df) != 0 {
		return df[0]
	}
	return v
}

func GetEnvInt(key string, defaultValue ...int) int {
	v := os.Getenv(key)
	if len(v) == 0 {
		return func() int {
			if len(defaultValue) != 0 {
				return defaultValue[0]
			}
			return 0
		}()
	}
	i, _ := strconv.Atoi(v)
	return i
}
