# SHA256 Hashes
```go
// [_SHA256 hash_](https://en.wikipedia.org/wiki/SHA-2) တွေကို
// binary သို့မဟုတ် စာသား blob တွေအတွက် တိုတောင်းတဲ့ identity တွေ တွက်ချက်ဖို့
// မကြာခဏ အသုံးပြုကြပါတယ်။ ဥပမာ၊ TLS/SSL လက်မှတ်တွေက SHA256 ကို
// လက်မှတ်တစ်ခုရဲ့ လက်မှတ်ထိုးချက်ကို တွက်ချက်ဖို့ သုံးပါတယ်။ Go မှာ SHA256 hash တွေကို
// ဘယ်လို တွက်ချက်မလဲဆိုတာ ဒီမှာရှင်းပြပါမယ်။

package main

// Go က hash function အမျိုးမျိုးကို `crypto/*` package တွေမှာ အကောင်အထည်ဖော်ထားပါတယ်။
import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"

	// ဒီမှာ ကျွန်တော်တို့က hash အသစ်တစ်ခုနဲ့ စပါတယ်။
	h := sha256.New()

	// `Write` က byte တွေကို မျှော်လင့်ပါတယ်။ သင့်မှာ string `s` ရှိရင်
	// `[]byte(s)` ကို သုံးပြီး byte တွေအဖြစ် ပြောင်းပါ။
	h.Write([]byte(s))

	// ဒါက နောက်ဆုံးရလဒ်ဖြစ်တဲ့ hash ကို byte slice အဖြစ် ရယူပါတယ်။
	// `Sum` ရဲ့ argument ကို ရှိပြီးသား byte slice တစ်ခုနဲ့ ပေါင်းဖို့ သုံးနိုင်ပါတယ်။
	// အဲဒါကို ပုံမှန်အားဖြင့် မလိုအပ်ပါဘူး။
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
```
```sh
# ပရိုဂရမ်ကို run လိုက်တာက hash ကို
#  တွက်ချက်ပြီး လူဖတ်လို့ရတဲ့ hex ပုံစံနဲ့ ပရင့်ထုတ်ပါတယ်။
$ go run sha256-hashes.go
sha256 this string
1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a...

# အထက်မှာပြထားတဲ့ ပုံစံနဲ့ ဆင်တူတဲ့ပုံစံကို
#  သုံးပြီး တခြား hash တွေကိုလည်း တွက်ချက်နိုင်ပါတယ်။
# ဥပမာ၊ SHA512 hash တွေ တွက်ချက်ဖို့ 
# `crypto/sha512` ကို import လုပ်ပြီး
# `sha512.New()` ကို သုံးလို့ရတယ်။

# သတိပြုရမှာက cryptographically 
# လုံခြုံတဲ့ hash တွေ လိုအပ်ရင်
# သင့်အနေနဲ့ [hash အင်အား](https://en.wikipedia.org/wiki/Cryptographic_hash_function) 
# အကြောင်းကို
# သေချာစွာ သုတေသနပြုလုပ်သင့်ပါတယ်!
```
