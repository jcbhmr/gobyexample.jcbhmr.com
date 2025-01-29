# Range over Channels->채널에서의 Range
```go
// [이전](/gobyexample/range) 예제에서 우리는 `for`와 `range`가 기본 자료구조에서 어떻게 순회하는지 살펴보았습니다.

package main

import "fmt"

func main() {

	// `queue` 채널에 있는 2개의 값을 순회 해보겠습니다.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// 이 `range`는 `queue`로부터 수신된 각 원소들을 순회합니다.
	//  위에서 채널을 `종료(close)`했기 때문에 이 순회 루프는 2개의 원소를 받으면 종료됩니다.
	for elem := range queue {
		fmt.Println(elem)
	}
}
```
```sh
$ go run range-over-channels.go
one
two

# 이 예제 또한 비어있지 않은 채널을 닫아도 채널에 남아있는 값들은 수신할 수 있다는걸 보여주고 있습니다.
```
