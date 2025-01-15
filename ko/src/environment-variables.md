# Environment Variables->환경변수
```go
// [환경변수](http://en.wikipedia.org/wiki/Environment_variable)
// 는 [Unix 프로그램에 설정 관련 정보를 전달](http://www.12factor.net/config)
// 하는데 일반적으로 사용되는 방법입니다.
//  환경변수를 어떻게 설정하는지, 어떻게 가져오는지, 그리고
//  설정된 환경변수들을 어떻게 보는지 봅시다.

package main

import "os"
import "strings"
import "fmt"

func main() {

	// 환경변수의 키-값 쌍을 설정하려면 `os.Setenv`를 이용합니다.
	//  어떤 키에 매핑된 값은 `os.Getenv`를 통해 가져올 수 있습니다.
	//  이 함수에 주어진 키가 환경변수로 설정되지 않은 키라면
	//  비어있는 문자열을 반환합니다.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// `os.Environ` 함수를 이용하여 환경변수로 설정된 모든
	//  키-값 쌍의 리스트를 볼 수 있습니다.
	//  이 함수는 `KEY=value`와 같은 형태의 문자열로 구성된
	//  슬라이스를 반환합니다. 슬라이스의 각 원소에
	//  `string.Split`과 같은 함수를 이용하여
	//  키와 값을 분리할 수 있습니다. 아래는 모든 키를 출력하는 예제입니다.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
```
```sh
# 프로그램을 돌려보면
#  `FOO`의 경우 우리가 프로그램 안에서 설정한 값을 가져오는 반면,
#  `BAR`의 경우는 빈 문자열을 가져오는 것을 볼 수 있습니다.
$ go run environment-variables.go
FOO: 1
BAR: 

# 환경변수로 지정된 키의 리스트는 컴퓨터에 따라 다르게 보일 수 있습니다.
TERM_PROGRAM
PATH
SHELL
...

# 우리가 환경변수로 `BAR`를 먼저 설정해두면,
#  실행될 프로그램이 해당 키에 설정된 값을 가져갑니다.
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...
```
