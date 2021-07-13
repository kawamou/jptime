package jptime

import (
	"time"
)

type Jptime struct {
	time.Time
}

func Wrap(t time.Time) Jptime {
	return Jptime{t}
}

// Compare はtimeパッケージのBefore/Afterのラッパーです
// 不等号でどちらの時間が前か（後か）判断します
// < > <= >= == の5種類を期待します
func (j Jptime) Compare(operator string, t time.Time) bool {
	switch operator {
	case "<":
		return j.Time.Before(t)
	case ">":
		return j.Time.After(t)
	case "<=":
		return !j.Time.After(t)
	case ">=":
		return !j.Time.Before(t)
	case "==":
		return j.Time.Equal(t)
	}
	return false
}

// IsHoliday は祝日かどうかを判定します
func (j Jptime) IsHoliday() bool {
	for _, holiday := range holidays {
		if isSameDay(j.Time, holiday.Date) {
			return true
		}
	}
	return false
}

// AllHolidays は祝日の一覧を返します
func (j Jptime) AllHolidays() []Holiday {
	return holidays
}

// isSameDay はふたつの日付が同日かを判定します
func isSameDay(t1, t2 time.Time) bool {
	return (t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day())
}
