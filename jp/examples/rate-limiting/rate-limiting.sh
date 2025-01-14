# プログラムを実行すると、はじめに送ったリクエストの一団は期待した通り200ミリ秒ごとに処理されたことがわかる。
$ go run rate-limiting.go
request 1 2012-10-19 00:38:18.687438 +0000 UTC
request 2 2012-10-19 00:38:18.887471 +0000 UTC
request 3 2012-10-19 00:38:19.087238 +0000 UTC
request 4 2012-10-19 00:38:19.287338 +0000 UTC
request 5 2012-10-19 00:38:19.487331 +0000 UTC

# その後に送られるリクエストのうち、はじめの3つはすぐに処理される。
# これはレート制限がバーストを許容するからだ。
# 最後に、残りの2つのリクエストが200ミリ秒ごとに処理される。
request 1 2012-10-19 00:38:20.487578 +0000 UTC
request 2 2012-10-19 00:38:20.487645 +0000 UTC
request 3 2012-10-19 00:38:20.487676 +0000 UTC
request 4 2012-10-19 00:38:20.687483 +0000 UTC
request 5 2012-10-19 00:38:20.887542 +0000 UTC
