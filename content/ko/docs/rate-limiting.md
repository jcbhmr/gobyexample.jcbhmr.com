# Rate Limiting->속도 제한
```go
// _[속도 제한(Rate limiting)](http://en.wikipedia.org/wiki/Rate_limiting)_은 리소스 이용을 제어하고 서비스의 품질을 유지하기위한 중요한 메커니즘입니다.
//  Go는 고루틴, 채널 그리고 [티커(tickers)](/gobyexample/tickers)로 속도 제한을 지원합니다.

package main

import "time"
import "fmt"

func main() {

	// 우선, 기본적인 속도 제한을 살펴봅시다. 요청 핸들링을 제한하고 싶다고 가정해봅시다.
	//  requests라는 이름으로 요청을 처리하는 채널을 만듭니다.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// 'limiter' 채널은 매 200 밀리초마다 값을 받습니다. 이게 우리 속도 제한 조절기입니다.
	limiter := time.Tick(time.Millisecond * 200)

	// 각 요청을 처리하기 전에 `limiter` 채널의 수신으로 블로킹 함으로써 매 200 밀리초마다 하나의 요청을 받도록 제한하고 있습니다.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 전반적으로는 속도 제한을 유지하면서 요청들을 짧게 버스팅(bursts of requests)하고 싶은 경우가 있습니다.
	//  limiter 채널을 버퍼링함으로써 이를 가능케할 수 있습니다.
	//  `burstyLimiter` 채널로 최대 3개의 이벤트를 버스팅할 수 있습니다.
	burstyLimiter := make(chan time.Time, 3)

	// 허용된 버스팅 수만큼 채널을 채웁니다.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 매 200 밀리초마다 최대 3개까지 `burstyLimiter`에 새로운 값을 추가합니다.
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	// 5개의 요청이 들어오는 시뮬레이션을 해봅시다. 처음 3개의 요청은 `burstyLimiter`의 버스팅 기능의 이점을 얻습니다.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
```
```sh
# 프로그램을 실행하면 첫번째 요청 배치들이 우리가 원하는대로 매 약 200 밀리초마다 처리되고 있음을 볼 수 있습니다.
$ go run rate-limiting.go
request 1 2012-10-19 00:38:18.687438 +0000 UTC
request 2 2012-10-19 00:38:18.887471 +0000 UTC
request 3 2012-10-19 00:38:19.087238 +0000 UTC
request 4 2012-10-19 00:38:19.287338 +0000 UTC
request 5 2012-10-19 00:38:19.487331 +0000 UTC

# 두번째 요청 배치에선 첫 3개의 요청은 버스팅 가능한 속도 제한덕에 즉각적으로 처리되며 나머지 2개의 요청은 각각 약 200 밀리초의 딜레이를 가지고 처리됩니다.
request 1 2012-10-19 00:38:20.487578 +0000 UTC
request 2 2012-10-19 00:38:20.487645 +0000 UTC
request 3 2012-10-19 00:38:20.487676 +0000 UTC
request 4 2012-10-19 00:38:20.687483 +0000 UTC
request 5 2012-10-19 00:38:20.887542 +0000 UTC
```
