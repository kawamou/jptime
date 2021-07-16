## Jptime
time.Timeのラッパーです。日本の祝日を扱えるようにしました。
### 使い方
```foo.go
t := time.Now()
// 祝日判定
if jptime.Wrap(t).IsHoliday() {
    // 何かしらの処理
}
// time.Timeのメソッドも全て使用可能です
// 例としてAddの場合
jptime.Wrap(t).Add(1 * time.Second)
```
### 祝日ファイル再生成
祝日ファイルは定期的にアップデートする必要があります。
祝日ファイル（`holidays.go`）をコマンドで再生成できるようにしました。

- `cmd/http_generate`ディレクトリ内で`go generate`コマンドを実行
    - 『内閣府「国民の祝日」について』ページへクローリングを行います
    - クローリングした`syukujitsu.csv`をパースし`holidays.go`を生成
- `cmd/local_generate`ディレクトリ内で`ARG=[filePath] go generate`コマンドを実行
    - `filePath`上にダウンロード済みの`syukujitsu.csv`をもとに`holidays.go`を生成