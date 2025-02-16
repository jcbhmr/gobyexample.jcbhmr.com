# SHA1 Hashes->SHA1 해시
```go
// [_SHA1 해시_(_SHA1 hashes_)](http://en.wikipedia.org/wiki/SHA-1)는 바이너리나 텍스트 블롭에 대한 짧은 식별자를 계산하기 위해 자주 사용됩니다.
//  예를 들어, [git revision control system](http://git-scm.com/)은 버저닝된 파일과 디렉토리를 식별하기 위해 광범위하게 SHA1을 사용합니다.
//  여기에 Go에서 SHA1 해시를 계산하는 방법이 있습니다.

package main

// Go는 다양한 `crypto/*` 패키지들에서 여러가지 해시 함수들을 구현하고 있습니다.
import "crypto/sha1"
import "fmt"

func main() {
	s := "sha1 this string"

	// 해시를 생성하는 패턴은 `sha1.New()`, `sha1.Write(bytes)`, 그리고 `sha1.Sum([]byte{})`입니다.
	//  새로운 해시를 만드는것부터 시작해봅시다.
	h := sha1.New()

	// `Write`는 바이트를 받습니다. 문자열을 `s`를 가지고 있다면, `[]byte(s)`를 사용해 이를 바이트로 강제하세요.
	h.Write([]byte(s))

	// 다음은 최종 해시 결괏값을 바이트 슬라이스로 얻습니다.
	//  `Sum`의 인자는 기존 바이트 슬라이스에 덧불이기 위해 사용될 수 있습니다: 보통은 필요하지 않습니다.
	bs := h.Sum(nil)

	// SHA1 값은 종종 16진수로 출력되는데, 예로 git 커밋이 있습니다.
	//  해쉬값을 16진수 문자열로 변환하기 위해선 `%x` 포맷 verb를 사용하세요.
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
```
```sh
# 프로그램을 실행하면 해시를 계산하고 이를 사람이 읽을 수 있는 16진수 포맷으로 출력합니다. 
$ go run sha1-hashes.go
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94


# 위에서 봤던 패턴과 유사한 패턴을 사용하여 다른 해시를 계산할 수 있습니다.
#  예를 들면, MD5 해시를 계산하기 위해선 `crypto/md5`를 임포트하고 `md5.New()`를 사용합니다.

# 참고로 암호학적으로 안전한 해시가 필요하다면 [해시 강도(hash strength)](http://en.wikipedia.org/wiki/Cryptographic_hash_function)를 주의깊게 조사해야합니다.
```
