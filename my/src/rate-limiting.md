# Rate Limiting
```go
// [_Rate limiting_](https://en.wikipedia.org/wiki/Rate_limiting) ဆိုတာ
// resource အသုံးပြုမှုကို ထိန်းချုပ်ဖို့နဲ့ ဝန်ဆောင်မှုအရည်အသွေးကို ထိန်းသိမ်းဖို့
// အရေးကြီးတဲ့ နည်းလမ်းတစ်ခု ဖြစ်ပါတယ်။ Go က goroutine တွေ၊
// channel တွေ၊ နဲ့ [ticker](tickers) တွေကိုသုံးပြီး rate limiting ကို
// လှပစွာ ပံ့ပိုးပေးပါတယ်။
package main

import (
	"fmt"
	"time"
)

func main() {

	// ပထမဆုံး rate limiting အခြေခံ ကို ကြည့်ကြပါမယ်။ ဝင်လာတဲ့ request တွေကို
	// ကိုင်တွယ်တာကို ကန့်သတ်ချင်တယ် ဆိုကြပါစို့။ ဒီ request တွေကို တူညီတဲ့နာမည်ရှိတဲ့
	// channel တစ်ခုကနေ ဖြန့်ဝေပေးပါမယ်။
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// ဒီ `limiter` channel က 200 မီလီစက္ကန့်တိုင်း တန်ဖိုးတစ်ခု လက်ခံရရှိပါမယ်။
	// ဒါဟာ ကျွန်တော်တို့ရဲ့ rate limiting scheme မှာ regulator အဖြစ် ဆောင်ရွက်ပါတယ်။
	limiter := time.Tick(200 * time.Millisecond)

	// Request တိုင်းကို မဖြန့်ဝေခင် `limiter` channel ကနေ လက်ခံတာကို block လုပ်ခြင်းအားဖြင့်
	// ကျွန်တော်တို့ဟာ 200 မီလီစက္ကန့်တိုင်း request တစ်ခုသာ လုပ်ဆောင်နိုင်အောင် ကန့်သတ်ထားပါတယ်။
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// ကျွန်တော်တို့ရဲ့ rate limiting scheme မှာ ပျမ်းမျှ rate limit ကို ထိန်းသိမ်းထားပြီး
	// တိုတောင်းတဲ့ request burst တွေကို ခွင့်ပြုချင်ပါတယ်။ ဒါကို limiter channel ကို
	// buffer လုပ်ခြင်းဖြင့် လုပ်ဆောင်နိုင်ပါတယ်။ ဒီ `burstyLimiter` channel က
	// 3 ခုအထိ burst ဖြစ်ခွင့်ပြုပါမယ်။
	burstyLimiter := make(chan time.Time, 3)

	// Bursting ခွင့်ပြုထားတာကို ကိုယ်စားပြုဖို့ channel ကို fill up လုပ်ပါမယ်။
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 200 မီလီစက္ကန့်တိုင်း `burstyLimiter` ထဲကို တန်ဖိုးအသစ်တစ်ခု ထည့်ဖို့ ကြိုးစားပါမယ်၊
	// သူ့ရဲ့ အများဆုံး 3 ခုဆိုတဲ့ ကန့်သတ်ချက်အထိပါ။
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// အခု နောက်ထပ် request 5 ခု ဝင်လာတာကို simulate လုပ်ပါမယ်။ ပထမ
	// 3 ခုက `burstyLimiter` ရဲ့ burst လုပ်နိုင်စွမ်းကနေ အကျိုးကျေးဇူး ရရှိပါလိမ့်မယ်။
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
# ကျွန်တော်တို့ရဲ့ ပရိုဂရမ်ကို run လိုက်တဲ့အခါ 
# ပထမအသုတ် request တွေကို
# ကျွန်တော်တို့ လိုချင်သလိုပဲ တစ်ခုနဲ့တစ်ခုကြား 
# ~200 မီလီစက္ကန့်စီခြားပြီး handle လုပ်တာ တွေ့ရပါတယ်။
$ go run rate-limiting.go
request 1 2012-10-19 00:38:18.687438 +0000 UTC
request 2 2012-10-19 00:38:18.887471 +0000 UTC
request 3 2012-10-19 00:38:19.087238 +0000 UTC
request 4 2012-10-19 00:38:19.287338 +0000 UTC
request 5 2012-10-19 00:38:19.487331 +0000 UTC

# ဒုတိယအသုတ် request တွေအတွက်တော့ 
# ပထမဆုံး request သုံးခုကို
# burstable rate limiting ကြောင့် 
# ချက်ချင်း ဖြေရှင်းပေးပါတယ်။
# ပြီးတော့မှ ကျန်တဲ့ နှစ်ခုကို တစ်ခုနဲ့တစ်ခုကြား 
# ~200ms စီ နှောင့်နှေးပြီး ဖြေရှင်းပေးပါတယ်။
request 1 2012-10-19 00:38:20.487578 +0000 UTC
request 2 2012-10-19 00:38:20.487645 +0000 UTC
request 3 2012-10-19 00:38:20.487676 +0000 UTC
request 4 2012-10-19 00:38:20.687483 +0000 UTC
request 5 2012-10-19 00:38:20.887542 +0000 UTC
```
