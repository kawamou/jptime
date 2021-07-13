package jptime

import (
	"encoding/csv"
	"io"
	"time"
)

// 『内閣府「国民の祝日」について』からsyukujitsu.csvを取得します
// https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv
// csvの形式はヘッダに続き祝日の日付(yyyy/mm/dd)と祝日の名称が並びます
// 下記に例を示します
// 国民の祝日・休日月日,国民の祝日・休日名称
// 1955/1/1,元日
// 1955/1/15,成人の日
// 1955/3/21,春分の日
// 1955/4/29,天皇誕生日
// 1955/5/3,憲法記念日
// 1955/5/5,こどもの日
// ...

// column はcsvが何列目かを表します
type column int

var (
	date column = 0 // 1列目には祝日のyyyy/mm/dd形式の日付が入ります
	name column = 1 // 2列目には祝日の名称（ex. 成人の日）が入ります
)

// layout はcsvの日付のレイアウトを表します
var layout = "2006/1/2"

// isHeader はsyukujitsu.csvのヘッダーを判定しスキップします
func isHeader(line []string) bool {
	return line[date] == "国民の祝日・休日月日" && line[name] == "国民の祝日・休日名称"
}

// Parse はsyukujitsu.csvを解析します
func Parse(r *csv.Reader) (holidays []Holiday, err error) {
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if isHeader(line) {
			continue
		}
		t, err := time.ParseInLocation(layout, line[date], time.Local)
		if err != nil {
			return nil, err
		}
		// TODO: len(line)が2に満たなかったときのケア
		holidays = append(holidays, Holiday{Date: t, Name: line[name]})
	}
	return holidays, nil
}
