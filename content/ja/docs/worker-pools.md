# Worker Pools
```go
// この例ではゴルーチンとチャネルを使って<em>ワーカープール</em>を実装するやり方を紹介する。

package main

import (
	"fmt"
	"time"
)

// このワーカーをいくつか平行に動かす。
// 各ワーカーは `jobs` チャネルに仕事を受信し、その結果を `results` チャネルに送信する。
// 実行に時間のかかる処理をシミュレートするために、仕事をするたびに1秒スリープする。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// ワーカープールを使うために、仕事を送り、結果を収集する必要がある。
	// そのために使う2つのチャネルを作る。
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// ここで3つワーカーを起動する。
	// まだジョブがないので、いずれのワーカーもブロックする。
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// `jobs` に値を5つ送り、その後チャネルを閉じて仕事を送り終えたことを伝える。
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 最後に仕事の結果を集める。
	// ここではゴルーチンのワーカーが終了したことも保証される。
	// 複数のゴルーチンが終了するのを待つには [WaitGroup](waitgroups) を使うやり方もある。
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
```
```sh
# プログラムを実行すると5つの仕事が色々なワーカーに実行される様子を確認できる。
# 仕事は合わせて5秒かかるはずだが、プログラムは2秒ほどしか使わない。
# これは、平行に処理を行うワーカーが3ついるためである。
$ time go run worker-pools.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real	0m2.358s
```
