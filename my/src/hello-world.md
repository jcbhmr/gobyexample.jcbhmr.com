# Hello World
```go
// ကျွန်ုပ်တို့၏ ပထမဆုံး print ထုတ်မယ့်ပရိုဂရမ်သည် ဂန္ထဝင် "Hello world" message ဘဲဖြစ်ပါတယ်။
// ဒီမှာ ကုဒ်အပြည့်အစုံကိုလည်းဖတ်နိုင်ပါတယ်။
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```
```sh
# ဒီပရိုဂရမ်ကိုစဖို့ရာအတွက် ဖိုင်ထဲမှာ
# ကုဒ်များကိုထည့်ပေးပါ။
# ပြီးသွားလျှင် `go run` ကိုသုံးပီး run လုပ်ပါ
$ go run hello-world.go
hello world

# တခါတလေ ကျနော်တို့ကိုယ့် program ကို
# binary အဖြစ် build လုပ်ပြီး
# run လုပ်ပြီးရှိလျှင် အဲ့အခါ `go build` ကိုသုံးပါ။
$ go build hello-world.go
$ ls
hello-world	hello-world.go

# binary build လုပ်ထားတဲ့ program
# ကိုဒီလိုတိုက်ရိုက် execute လုပ်နိုင်ပါတယ်
$ ./hello-world
hello world

# ကျနော်တို့လက်ရှိ Go program ကို run
# နှင့် build လုပ်နိုင်ပြီးရှိလျှင်
# ကျနော်တို့လက်ရှိ Go language ကိုအသေးစိတ်လေ့လာကြရအောင်။
```
