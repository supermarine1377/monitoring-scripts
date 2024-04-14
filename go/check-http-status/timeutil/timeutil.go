package timeutil

import "time"

const timeFormat = "2006-01-02_15-04-05"

func NowStr() string {
	return time.Now().Local().Format(timeFormat)
}
