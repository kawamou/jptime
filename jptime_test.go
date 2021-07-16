package jptime

import (
	"testing"
	"time"
)

func TestJptime_IsHoliday(t *testing.T) {
	tests := []struct {
		name string
		arg  time.Time
		want bool
	}{
		{
			name: "2021年7月19日",
			arg:  time.Date(2021, 7, 19, 0, 0, 0, 0, time.Local),
			want: false,
		},
		{
			name: "2021年7月20日",
			arg:  time.Date(2021, 7, 20, 0, 0, 0, 0, time.Local),
			want: false,
		},
		{
			name: "2021年7月21日",
			arg:  time.Date(2021, 7, 21, 0, 0, 0, 0, time.Local),
			want: false,
		},
		{
			name: "2021年7月22日",
			arg:  time.Date(2021, 7, 22, 0, 0, 0, 0, time.Local),
			want: true,
		},
		{
			name: "2021年7月23日",
			arg:  time.Date(2021, 7, 23, 0, 0, 0, 0, time.Local),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !(Wrap(tt.arg).IsHoliday() == tt.want) {
				t.Fail()
			}
		})
	}
}
