# Goroutines
```go
// _goroutine_ ဆိုတာ ပေါ့ပါးတဲ့ execution thread တစ်ခုဖြစ်ပါတယ်။

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// `f(s)` ဆိုတဲ့ function ရှိတယ်ထားပါတော့ ဒီမှာ ပုံမှန်အတိုင်း
	// synchronous(တပြိုင်နက်တည်း) ခေါ်တာကို ပြထားပါတယ်။
	f("direct")

	// ဒီ function ကို goroutine အနေနဲ့ ခေါ်ချင်ရင် `go f(s)` လို့သုံးပါတယ်။
	// ဒီ goroutine အသစ်က ခေါ်လိုက်တဲ့ goroutine နဲ့ တပြိုင်နက်တည်း (concurrently) အလုပ်လုပ်ပါလိမ့်မယ်။
	go f("goroutine")

	// Anonymous function ကိုလည်း goroutine အနေနဲ့ စတင်နိုင်ပါတယ်။
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// အခု ကျွန်တော်တို့ရဲ့ function ခေါ်တာ နှစ်ခုက သီးခြား goroutine တွေမှာ asynchronous (တပြိုင်နက်တည်းမဟုတ်) အနေနဲ့ အလုပ်လုပ်နေပါပြီ။
	// သူတို့ပြီးဆုံးတာကို စောင့်နေတာပါ (ပို robust approach တွေအတွက်ကတော့ [WaitGroup](waitgroups) ကို သုံးပါ)။
	time.Sleep(time.Second)
	fmt.Println("done")
}
```
```sh
# ဒီပရိုဂရမ်ကို run လိုက်တဲ့အခါ၊ ကျွန်တော်တို့က 
# blocking call ရဲ့ output ကို အရင်မြင်ရပြီး
# နောက်မှ goroutine နှစ်ခုရဲ့ output ကို မြင်ရပါတယ်။
# goroutine တွေရဲ့ output က
# ရောနှောနေနိုင်ပါတယ်။ ဘာကြောင့်လဲဆိုတော့ 
# goroutine တွေက Go runtime က
# တပြိုင်နက်တည်း concurrent ဖြစ်အောင် run နေလို့ပါ။.
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done

# နောက်တစ်ဆင့်မှာ Go ရဲ့ concurrent 
# ပရိုဂရမ်တွေမှာ goroutine တွေနဲ့အတူ
# တွဲဖက်သုံးလေ့ရှိတဲ့ channel တွေအကြောင်း
#  လေ့လာကြပါမယ်။
```
