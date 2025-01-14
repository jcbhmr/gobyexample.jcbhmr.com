// Go ではパターンに基づくレイアウトを利用して時刻をフォーマットできる。

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// まずは RFC3339 に対応するレイアウトを使って時刻をフォーマットする方法を紹介する。
	t := time.Now()
	p(t.Format(time.RFC3339))

	// 時刻をパースするときも `Format` と同じレイアウトを使う。
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	// `Format`、`Parse` では例示に基づいてレイアウトを決める。
	// ふつうは `time` モジュールに定義されている定数をレイアウトの例として使うが、独自のレイアウトを使ってもよい。
	// レイアウトでは特定の時刻 `Mon Jan 2 15:04:05 MST 2006` を表している必要があり、プログラムはこれに従って時刻をフォーマットしたり、文字列をパースしたりする。
	// 時刻の例はちょうど以下のようなものだ。例えば、以下の例では、年は2006に、時間は15に、曜日は月曜になっている。
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	// フォーマットしたいのが数値なら、時刻値の部分を抜き出し、文字列のフォーマット機能を使ってもよい。
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// `Parse` は入力が不正だと、パース時に起きた問題を説明するエラーを返す。
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}