# Worker Pools->워커 풀
```go
// 이번 예제에서는 고루틴과 채널로 _워커풀(worker pool)을 구현하는 방법을 살펴보겠습니다.

package main

import "fmt"
import "time"

// 다음은 여러개의 인스턴스를 동시에 실행시킬 워커입니다.
//  이 워커들은 `jobs` 채널을 통해 작업을 받으며 작업의 결괏값을 `results`로 보냅니다.
//  비용이 큰 작업을 시뮬레이션 하기 위해 각 잡마다 1초의 딜레이를 줄겁니다.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// 워커풀을 사용하기 위해선 워커에 작업을 보내고 그 결괏값들을 수집해야합니다.
	//  이를 위해 2개의 채널을 만듭니다.
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 다음은 3개의 워커를 실행시키는데, 처음에는 잡이 없기 때문에 각 워커는 블로킹 됩니다.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 5개의 잡을 보내고 작업을 다 보냈음을 알리기위해 채널을 `close`합니다.
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// 마지막으로 모든 작업의 결괏값들을 가져옵니다.
	for a := 1; a <= 5; a++ {
		<-results
	}
}
```
```sh
# 프로그램은 5개의 잡이 여러개의 워커에 의해 실행되고 있음을 보여줍니다.
#  총 작업 시간은 5초지만 3개의 워커가 동시에 실행되고 있기 때문에, 전체 작업은 약 2초만에 끝납니다.
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
