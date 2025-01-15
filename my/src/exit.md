# Exit
```go
// အခြေအနေ (status) တစ်ခုနဲ့ ချက်ချင်းထွက်ဖို့ `os.Exit` ကို သုံးပါ။

package main

import (
	"fmt"
	"os"
)

func main() {

	// `defer` တွေကို `os.Exit` သုံးတဲ့အခါ run မှာ မဟုတ်ပါဘူး၊ ဒါကြောင့်
	// ဒီ `fmt.Println` ကို ဘယ်တော့မှ ခေါ်မှာ မဟုတ်ပါဘူး။
	defer fmt.Println("!")

	// အခြေအနေ(status) 3 နဲ့ ထွက်ပါ။
	os.Exit(3)
}

// C စတဲ့ အခြားဘာသာစကားတွေနဲ့ မတူဘဲ၊ Go က `main` ကနေ integer တန်ဖိုး
// ပြန်ပေးတာကို exit status အဖြစ် မသုံးပါဘူး။ သင်က သုည မဟုတ်တဲ့
// အခြေအနေ(status)နဲ့ ထွက်ချင်ရင် `os.Exit` ကို သုံးသင့်ပါတယ်။
```
```sh
# သင်က `exit.go` ကို `go run` နဲ့ run ရင်၊ exit က
# `go` က ဖမ်းယူပြီး ပုံနှိပ်ဖော်ပြပေးပါလိမ့်မယ်။
$ go run exit.go
exit status 3
# Binary ကို build လုပ်ပြီး run ခြင်းဖြင့် သင်က
# terminal ထဲမှာ status ကို မြင်နိုင်ပါတယ်။
$ go build exit.go
$ ./exit
$ echo $?
3
# မှတ်ချက်။ ကျွန်ုပ်တို့ program ထဲက `!` 
# က ဘယ်တော့မှ ပုံနှိပ် (print) မထွက်လာပါ။
```
