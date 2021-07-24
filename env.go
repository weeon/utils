package utils

import "os"

func GetEvn(key string, df ...string) string {
	v := os.Getenv(key)
	if len(v) == 0 && len(df) != 0 {
		return df[0]
	}
	return v
}
