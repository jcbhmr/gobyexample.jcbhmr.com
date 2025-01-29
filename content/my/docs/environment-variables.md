# Environment Variables
```go
// [Environment variables](https://en.wikipedia.org/wiki/Environment_variable)
// သည် [Unix ပရိုဂရမ်များသို့ configuration အချက်အလက်များ ပေးပို့ရန်အတွက်](https://www.12factor.net/config)
// universal နည်းလမ်းတစ်ခုဖြစ်သည်။
// Environment variables များကို မည်သို့ သတ်မှတ်ခြင်း၊ ရယူခြင်းနှင့် စာရင်းပြုစုခြင်း(list)တို့ကို ကြည့်ကြပါစို့။

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// key/value တွဲတစ်ခုကို သတ်မှတ်ရန် `os.Setenv` ကို အသုံးပြုပါ။
	// key တစ်ခုအတွက် တန်ဖိုးကို ရယူရန် `os.Getenv` ကို အသုံးပြုပါ။
	// အကယ်၍ key သည် environment တွင် မရှိပါက ဤ function သည်
	// စာကြောင်းအလွတ်တစ်ခုကို ပြန်ပေးပါလိမ့်မည်။
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// environment ရှိ key/value တွဲအားလုံးကို စာရင်းပြုစု(list)ရန် `os.Environ` ကို အသုံးပြုပါ။
	// ဤ function သည် `KEY=value` ပုံစံဖြင့် စာကြောင်းများ၏ slice တစ်ခုကို ပြန်ပေးပါသည်။
	// သင်သည် ၎င်းတို့ကို `strings.SplitN` ဖြင့် key နှင့် value ကို ရရှိနိုင်ပါသည်။
	// ဤနေရာတွင် ကျွန်ုပ်တို့သည် key အားလုံးကို ပုံနှိပ်ဖော်ပြထားပါသည်။
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
```
```sh
# ပရိုဂရမ်ကို run ခြင်းဖြင့် ကျွန်ုပ်တို့သည် 
# ပရိုဂရမ်ထဲတွင် သတ်မှတ်ထားသော
# `FOO` ၏ တန်ဖိုးကို ရယူနိုင်ကြောင်း တွေ့ရပါသည်။ 
# သို့သော် 
# `BAR` သည် အလွတ်ဖြစ်နေပါသည်။
$ go run environment-variables.go
FOO: 1
BAR: 

# ပတ်ဝန်းကျင်ရှိ key များ၏ စာရင်းသည် 
# သင့်ကွန်ပျူတာပေါ်မူတည်၍ 
# ကွဲပြားနိုင်ပါသည်။
TERM_PROGRAM
PATH
SHELL
...
FOO

# အကယ်၍ ကျွန်ုပ်တို့သည် `BAR` ကို environment 
# တွင် ဦးစွာ သတ်မှတ်ပါက 
# run နေသော ပရိုဂရမ်သည် ထိုတန်ဖိုးကို ရယူပါလိမ့်မည်။
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...
```
