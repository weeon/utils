package utils

import "time"

const(
	DefaultTimeFormat = "2006-01-02 15:04:05"
)

func FormatTime(t time.Time) string {
	return t.Format(DefaultTimeFormat)
}

func TimeFromUnix(i int64) time.Time{
	return time.Unix(i,0)
}