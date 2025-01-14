# ちょうど50000回操作を実行するはずだ。
# このときにアトミックではない `ops++` を使うと、ゴルーチンが相互に干渉しあうので、
# カウンタの値は実行のたびに異なる（そして50000でもない）値になるだろう。
# `-race` フラグを使ってデータアクセスの競合を検出することもできる。
$ go run atomic-counters.go
ops: 50000

# 次の例では、状態管理のためのもう一つのツールである、
# ミューテックスを紹介する。