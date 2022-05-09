package timeutil

import "time"

const (
	DefaultTimeLayout = "2006-01-02 15:04:05"
	DayTimeLayout     = "2006-01-02 00:00:00"
)

// CurDay 今天的时间
func CurDay() string {
	return time.Now().Format(DayTimeLayout)
}

// Parse 转换标准time.Time
func Parse(t string) (time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.ParseInLocation(DefaultTimeLayout, t, loc)
}

// Before 计算t1是否在t2之前
func Before(t1, t2 string) (bool, error) {
	time1, err := Parse(t1)
	time2, err := Parse(t2)
	if err != nil {
		return false, err
	}
	return time1.Before(time2), nil
}
