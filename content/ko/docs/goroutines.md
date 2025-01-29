# Goroutines->고루틴
```go
// _goroutine_은 경량의 실행 스레드입니다.

package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// 함수 `f(s)`를 호출한다고 해봅시다. 다음은 동기적으로 실행되는 일반적인 방법으로 함수를 호출하는 방식을 보여줍니다.
	f("direct")

	// 이 함수를 고루틴으로 실행하기위해, `go f(s)`를 사용합니다.
	//  새로운 고루틴은 이를 호출한 루틴과 동시에 실행됩니다.
	go f("goroutine")

	// 익명 함수 호출에 대해서도 고루틴을 사용할 수 있습니다.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 우리가 호출한 두 개의 함수가 이제 별도의 고루틴에서 비동기적으로 실행되므로, 실행은 여기로 떨어집니다. (즉, 실행이 이 위치로 진행됨)
	//  `Scanln` 코드는 프로그램이 종료되기 전에 키 입력을 필요로 합니다.
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
```
```sh
# 이 프로그램을 실행하면, 블로킹된 호출(동기적으로 호출된 함수)의 결과가 먼저 출력된 후에, 인터리브된 두 개의 고루틴의 결괏값이 출력됩니다.
#  이 인터리빙은 Go 런타임에 의해 고루틴들이 동시에 실행되는 프로세스를 반영합니다.
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
<enter>
done

# 다음으로, Go 프로그램에서 동시에 실행되는 고루틴들을 보완해주는 채널(channels)에 대해 알아보겠습니다.
```
