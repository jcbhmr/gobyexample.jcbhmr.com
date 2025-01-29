# Random Numbers
```go
// Go ရဲ့ `math/rand` package က
// [pseudorandom number](https://en.wikipedia.org/wiki/Pseudorandom_number_generator)
// ထုတ်လုပ်မှုကို ပေးပါတယ်။

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// ဥပမာ၊ `rand.Intn` က `0 <= n < 100` ဖြစ်တဲ့
	// ကျပန်း `int` n ကို ပြန်ပေးပါတယ်။
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// `rand.Float64` က `0.0 <= f < 1.0` ဖြစ်တဲ့
	// `float64` `f` ကို ပြန်ပေးပါတယ်။
	fmt.Println(rand.Float64())

	// ဒါကို တခြား range တွေအတွက် ကျပန်း float တွေ ထုတ်လုပ်ဖို့ သုံးနိုင်ပါတယ်၊
	// ဥပမာ `5.0 <= f' < 10.0`။
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// ပုံမှန် number generator က deterministic ဖြစ်တဲ့အတွက် ပုံမှန်အားဖြင့်
	// အကြိမ်တိုင်းမှာ တူညီတဲ့ နံပါတ်စီးကြောင်းကို ထုတ်ပေးပါလိမ့်မယ်။
	// ကွဲပြားတဲ့ စီးကြောင်းတွေ ထုတ်လုပ်ဖို့ ပြောင်းလဲနေတဲ့ seed တစ်ခု ပေးပါ။
	// သတိပြုရမှာက လျှို့ဝှက်ထားချင်တဲ့ ကျပန်းနံပါတ်တွေအတွက် ဒါကို သုံးလို့ မလုံခြုံပါဘူး။
	// အဲဒီအတွက် `crypto/rand` ကို သုံးပါ။
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// ရလာတဲ့ `rand.Rand` ကို `rand` package ပေါ်က
	// function တွေလိုပဲ ခေါ်သုံးပါ။
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// တူညီတဲ့နံပါတ်နဲ့ source တစ်ခုကို seed လုပ်ရင်
	// တူညီတဲ့ ကျပန်းနံပါတ်စီးကြောင်းကို ထုတ်လုပ်ပါတယ်။
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
```
```sh
# ဒီနမူနာကို သင်ဘယ်နေရာမှာ run 
# သလဲပေါ်မူတည်ပြီး ထုတ်လုပ်ထားတဲ့ နံပါတ်တချို့
# ကွာခြားနိုင်ပါတယ်။ Go playground မှာတော့ 
# `time.Now()` နဲ့ seed လုပ်တာက
# playground ရဲ့ implementation ကြောင့် 
# အခြေအနေတူတဲ့ ရလဒ်တွေကိုပဲ
#  ထုတ်ပေးနေဦးမှာကို သတိပြုပါ။
$ go run random-numbers.go
81,87
0.6645600532184904
7.123187485356329,8.434115364335547
0,28
5,87
5,87

# Go က ပေးနိုင်တဲ့ တခြား ကျပန်းပမာဏတွေအတွက် 
# ကိုးကားချက်တွေကို
# [`math/rand`](https://pkg.go.dev/math/rand) package
#  docs မှာ ကြည့်ပါ။
```
