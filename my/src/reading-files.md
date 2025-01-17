# Reading Files
```go
// ဖိုင်များဖတ်ခြင်းနှင့် ရေးခြင်းသည် Go ပရိုဂရမ်များအတွက် အခြေခံလိုအပ်ချက်ဖြစ်သည်။
// ပထမဦးစွာ ဖိုင်ဖတ်ခြင်းဥပမာအချို့ကို ကြည့်ကြမည်။

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// ဖိုင်ဖတ်ခြင်းသည် အများအားဖြင့် အမှားစစ်ဆေးမှုလိုအပ်သည်။
// ဤ helper function သည် အောက်ပါ အမှားစစ်ဆေးမှုများကို ရိုးရှင်းစေမည်။
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// အခြေခံကျဆုံး ဖိုင်ဖတ်ခြင်းလုပ်ငန်းမှာ ဖိုင်တစ်ခုလုံး၏အကြောင်းအရာကို
	// မှတ်ဉာဏ်ထဲ (memory) သို့ တစ်ခါတည်းဖတ်ယူခြင်းဖြစ်သည်။
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// ဖိုင်၏မည်သည့်အပိုင်းကို မည်သို့ဖတ်မည်ကို ပိုမိုထိန်းချုပ်လိုကြောင်း သင်တွေ့ရလိမ့်မည်။
	// ဤလုပ်ငန်းများအတွက် 'Open' ဖြင့် ဖိုင်ကိုဖွင့်ပြီး 'os.File' တန်ဖိုးကို ရယူပါ။
	f, err := os.Open("/tmp/dat")
	check(err)

	// ဖိုင်၏အစမှ byte အချို့ကိုဖတ်ပါ။
	// 5 byte အထိဖတ်ခွင့်ပြုသော်လည်း အမှန်တကယ်ဖတ်မိသည့် အရေအတွက်ကိုလည်း မှတ်သားပါ။
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// သင်သည် ဖိုင်ထဲတွင် သိထားသောနေရာသို့ 'Seek' လုပ်နိုင်ပြီး ထိုနေရာမှ 'Read' လုပ်နိုင်သည်။
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// 'io' package သည် ဖိုင်ဖတ်ခြင်းအတွက် အသုံးဝင်သော function များပေးသည်။
	// ဥပမာအားဖြင့် အထက်ပါကဲ့သို့ ဖတ်ခြင်းများကို 'ReadAtLeast' ဖြင့် ပိုမိုခိုင်မာစွာ အကောင်အထည်ဖော်နိုင်သည်။
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// ပါဝင်ပြီးသား rewind မရှိသော်လည်း 'Seek(0, 0)' သည် ဤလုပ်ငန်းကို ဆောင်ရွက်နိုင်သည်။
	_, err = f.Seek(0, 0)
	check(err)

	// 'bufio' package သည် buffered reader ကို အကောင်အထည်ဖော်သည်။
	// ၎င်းသည် အသေးစား ဖတ်ခြင်းများစွာအတွက် ထိရောက်မှုရှိခြင်းနှင့်
	// ၎င်းပေးသော ထပ်ဆောင်းဖတ်ခြင်းနည်းလမ်းများကြောင့် အသုံးဝင်နိုင်သည်။
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// အသုံးပြုပြီးပါက ဖိုင်ကိုပိတ်ပါ (ပုံမှန်အားဖြင့် ၎င်းကို 'Open' လုပ်ပြီးချက်ချင်း
	// 'defer' ဖြင့် သတ်မှတ်လေ့ရှိသည်)။
	f.Close()
}
```
```sh
$ echo "hello" > /tmp/dat
$ echo "go" >>   /tmp/dat
$ go run reading-files.go
hello
go
5 bytes: hello
2 bytes @ 6: go
2 bytes @ 6: go
5 bytes: hello

# နောက်တစ်ဆင့်အနေဖြင့် 
# ဖိုင်ရေးသားခြင်းကို လေ့လာကြမည်။
```
