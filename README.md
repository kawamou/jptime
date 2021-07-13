# Jptime
- time.Timeのラッパーです
    - 日本の祝日を扱えるようにしました
    - time.TimeのBeforeやAfter等の比較が分かりづらかったのでCompareメソッドを作成しました
## 使い方
```foo.go
t := time.Now()
// 比較
if jptime.Wrap(t).Compare("<", time.Now()) {
    // 何かしらの処理
}
// 祝日判定
if jptime.Wrap(t).IsHoliday() {
    // 何かしらの処理
}
// time型に具備されたメソッドも全て使えます
jptime.Wrap(t).Add(1 * time.Second)
```
## 祝日ファイル生成
- 祝日ファイルが古くなる可能性も考え、祝日ファイル（holidays.go）を再生成できるようにしました
- cmd > http_generate内で`go generate`
    - 『内閣府「国民の祝日」について』ページへクローリングしてファイルを生成します
    - syukujitsu.csvというファイルをパースしGoファイルに変換しています
- cmd > local_generate内で`ARG=filePath go generate`
    - ローカルにあるsyukujitsu.csvをもとにファイルを生成します
## 用語定義
- Weekday：週日。祝日を含む月〜金
- Businessday：平日。祝日を除く週日
- Weekend：週末。土・日
- Holiday：祝日。国民の休日等を表す