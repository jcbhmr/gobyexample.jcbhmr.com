# Command-Line Arguments
```go
// [_Command-line arguments_](https://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// သည် ပရိုဂရမ်များ၏ လုပ်ဆောင်မှုကို parameter ပေးရန် အသုံးများသော နည်းလမ်းတစ်ခုဖြစ်သည်။
// ဥပမာ - `go run hello.go` သည် `go` ပရိုဂရမ်အတွက် `run` နှင့် `hello.go` arguments များကို အသုံးပြုသည်။

package main

import (
	"fmt"
	"os"
)

func main() {

	// `os.Args` သည် ကနဦး command-line arguments များကို ရယူခွင့်ပေးသည်။
	// ဤ slice ထဲရှိ ပထမဆုံးတန်ဖိုးသည် ပရိုဂရမ်၏ လမ်းကြောင်းဖြစ်ပြီး
	// `os.Args[1:]` သည် ပရိုဂရမ်သို့ ပေးသော arguments များကို ကိုင်တွယ်(hold) သည်ကို သတိပြုပါ။
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// သင်သည် ပုံမှန် indexing ဖြင့် တစ်ခုချင်းစီသော args များကို ရယူနိုင်သည်။
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
```
```sh
# Command-line arguments တွေကို စမ်းသပ်ဖို့အတွက် 
# ပထမဆုံး `go build` နဲ့ binary ကို 
# တည်ဆောက်တာ(build)က အကောင်းဆုံးဖြစ်ပါတယ်။
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]       
[a b c d]
c

# နောက်တစ်ဆင့်အနေနဲ့ ပိုပြီးအဆင့်မြင့်တဲ့ 
# command-line processing ကို
# flags တွေသုံးပြီး လေ့လာကြမှာ ဖြစ်ပါတယ်။
```
