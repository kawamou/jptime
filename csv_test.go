package jptime

import (
	"encoding/csv"
	"strings"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want []Holiday
	}{
		{
			name: "元日",
			arg:  `1955/1/1,元日`,
			want: []Holiday{
				{
					Date: time.Date(1955, 1, 1, 0, 0, 0, 0, time.Local),
					Name: "元日",
				},
			},
		},
		{
			name: "休日",
			arg:  `1993/5/4,休日`,
			want: []Holiday{
				{
					Date: time.Date(1993, 5, 4, 0, 0, 0, 0, time.Local),
					Name: "休日",
				},
			},
		},
		{
			name: "スポーツの日",
			arg:  `2019/10/14,体育の日（スポーツの日）`,
			want: []Holiday{
				{
					Date: time.Date(2019, 10, 14, 0, 0, 0, 0, time.Local),
					Name: "体育の日（スポーツの日）",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.arg)
			reader := csv.NewReader(r)
			parsed, err := Parse(reader)
			if err != nil {
				t.Log(err)
				t.Fail()
			}
			for _, p := range parsed {
				for _, w := range tt.want {
					if p == w {
						continue
					} else {
						t.Fail()
					}
				}
			}
		})
	}
}
