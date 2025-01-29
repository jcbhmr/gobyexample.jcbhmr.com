# Channels
```go
// _Channel_ တွေဆိုတာ concurrent ဖြစ်နေတဲ့ goroutine တွေကို ဆက်သွယ်ပေးတဲ့ ပိုက်လိုင်းတွေပါ။
// goroutine တစ်ခုကနေ channel ထဲကို တန်ဖိုးတွေ ပို့နိုင်ပြီး
// နောက် goroutine တစ်ခုကနေ အဲဒီတန်ဖိုးတွေကို လက်ခံယူနိုင်ပါတယ်။
package main

import "fmt"

func main() {

	// `make(chan val-type)` နဲ့ channel အသစ်တစ်ခု ဖန်တီးပါတယ်။
	// Channel တွေက သူတို့သယ်ဆောင်တဲ့ တန်ဖိုးတွေရဲ့ type အလိုက် type သတ်မှတ်ချက်ရှိပါတယ်။
	messages := make(chan string)

	// `channel <-` syntax ကိုသုံးပြီး channel ထဲကို တန်ဖိုးတစ်ခု _ပို့_ ပါတယ်။
	// ဒီမှာ အပေါ်မှာဖန်တီးထားတဲ့ `messages` channel ထဲကို `"ping"` ဆိုတဲ့စကားလုံးကို
	// goroutine အသစ်တစ်ခုကနေ ပို့လိုက်ပါတယ်။
	go func() { messages <- "ping" }()

	// `<-channel` syntax က channel ကနေ တန်ဖိုးတစ်ခုကို _လက်ခံယူ(receive)_ ပါတယ်။
	// ဒီမှာ အပေါ်မှာပို့လိုက်တဲ့ `"ping"` message ကို လက်ခံယူပြီး print ထုတ်ပါမယ်။
	msg := <-messages
	fmt.Println(msg)
}
```
```sh
# ပရိုဂရမ်ကို run လိုက်တဲ့အခါ `"ping"` message က 
# ကျွန်တော်တို့ channel ကနေတဆင့် goroutine
#  တစ်ခုကနေ နောက်တစ်ခုဆီကို 
# အောင်မြင်စွာ ပို့ဆောင်ခဲ့ပါတယ်။
$ go run channels.go 
ping

# ပုံမှန်အားဖြင့် ပို့တာနှင့်လက်ခံတာက 
# ပို့သူနဲ့လက်ခံသူ နှစ်ဦးစလုံး အဆင်သင့်ဖြစ်တဲ့အထိ 
# block လုပ်ထားပါတယ်။
# ဒီ property က ကျွန်တော်တို့ ပရိုဂရမ်အဆုံးမှာ 
# `"ping"` message အတွက်၊
# တခြား synchronization နည်းလမ်းတွေ 
# သုံးစရာမလိုပဲနဲ့ စောင့်ဆိုင်းနိုင်စေပါတယ်။
```
