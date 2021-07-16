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
