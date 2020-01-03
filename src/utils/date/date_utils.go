package date

import "time"

const (
	apiDateLayout = "2006-JAN-02T15:04:05"
	apiDbLayout   = "2006-JAN-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNowDbFormat() string {
	return GetNow().Format(apiDbLayout)
}
