# Command-Line Flags
```go
// [_Command-line flags_](https://en.wikipedia.org/wiki/Command-line_interface#Command-line_option)
// ကိုကွန်မန်းလိုင်းပရိုဂရမ်များအတွက် option များသတ်မှတ်ရန် အသုံးများသောနည်းလမ်းတစ်ခုဖြစ်သည်။
// ဥပမာ `wc -l` တွင် `-l` သည် command-line flag တစ်ခုဖြစ်သည်။

package main

// Go သည် အခြေခံ command-line flag parsing ကို ထောက်ပံ့ပေးသော `flag` package ကို ပေးထားသည်။
// ကျွန်ုပ်တို့သည် ဤ package ကို အသုံးပြု၍ ကျွန်ုပ်တို့၏ နမူနာ command-line program ကို အကောင်အထည်ဖော်မည်။
import (
	"flag"
	"fmt"
)

func main() {

	// အခြေခံ flag ကြေညာချက်များကို string, integer နှင့် boolean option များအတွက် ရရှိနိုင်သည်။
	// ဤနေရာတွင် ကျွန်ုပ်တို့သည် `"foo"` မူလတန်ဖိုးနှင့် တိုတောင်းသော ဖော်ပြချက်ပါသော `word` string flag ကို ကြေညာသည်။
	// ဤ `flag.String` function သည် string pointer တစ်ခု (string တန်ဖိုးမဟုတ်) ကို ပြန်ပေးသည်။
	// ကျွန်ုပ်တို့သည် ဤ pointer ကို မည်သို့အသုံးပြုမည်ကို အောက်မှာတွေ့ရပါ့မယ်။
	wordPtr := flag.String("word", "foo", "a string")

	// ဤသည်မှာ `word` flag နှင့် ဆင်တူသောနည်းလမ်းကို အသုံးပြု၍ `numb` နှင့် `fork` flag များကို ကြေညာခြင်းဖြစ်သည်။
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	// ပရိုဂရမ်၏ အခြားနေရာတွင် ကြေညာထားပြီးသား var ကို အသုံးပြုသော option ကိုလည်း ကြေညာနိုင်သည်။
	// flag ကြေညာခြင်း function သို့ pointer တစ်ခုကို ပေးပို့ရန် လိုအပ်ကြောင်း သတိပြုပါ။
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// flag အားလုံးကြေညာပြီးနောက် command-line parsing ကို အကောင်အထည်ဖော်ရန် `flag.Parse()` ကို ခေါ်ပါ။
	flag.Parse()

	// ဤနေရာတွင် ကျွန်ုပ်တို့သည် parse လုပ်ထားသော option များနှင့် နောက်ဆက်တွဲ positional argument များကို dump လုပ်မည်။
	// အမှန်တကယ် option တန်ဖိုးများကို ရရှိရန် `*wordPtr` စသည်ဖြင့် pointer များကို dereference လုပ်ရန် လိုအပ်ကြောင်း သတိပြုပါ။
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
```
```sh
# command-line flags ပရိုဂရမ်ကို 
# စမ်းသပ်ရန်အတွက်
# ပထမဦးစွာ compile လုပ်ပြီး ထွက်လာသော 
# binary ကို တိုက်ရိုက် run သင့်သည်။
$ go build command-line-flags.go

# flag အားလုံးအတွက် တန်ဖိုးများပေး၍ 
# build လုပ်ထားသော ပရိုဂရမ်ကို စမ်းသပ်ကြည့်ပါ။
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# flag များကို ချန်လှပ်ထားပါက ၎င်းတို့သည် 
# အလိုအလျောက် ၎င်းတို့၏ မူလ(default) တန်ဖိုးများကို 
# ယူကြောင်း သတိပြုပါ။
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# နောက်ဆက်တွဲ positional argument 
# များကို flag များ၏နောက်တွင် ပေးနိုင်သည်။
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# `flag` package သည် flag အားလုံးကို positional
#  argument များ၏ရှေ့တွင် ပေါ်လာရန် 
# လိုအပ်ကြောင်း သတိပြုပါ 
# (သို့မဟုတ်ပါက flag များကို positional argument 
# များအဖြစ် အဓိပ္ပာယ်ကောက်ယူမည်)။
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# command-line ပရိုဂရမ်အတွက် အလိုအလျောက်ထုတ်ပေးသော 
# help စာသားကို ရရှိရန် `-h` သို့မဟုတ် `--help` 
# flag များကို အသုံးပြုပါ။
$ ./command-line-flags -h
Usage of ./command-line-flags:
 -fork=false: a bool
 -numb=42: an int
 -svar="bar": a string var
 -word="foo": a string

# `flag` package တွင် သတ်မှတ်မထားသော flag 
# တစ်ခုကို ပေးပါက ပရိုဂရမ်သည် အမှားပြ message 
# ကို ပုံနှိပ်ထုတ်ပေးပြီး help စာသားကို ထပ်ပြပေးမည်။
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...
```
