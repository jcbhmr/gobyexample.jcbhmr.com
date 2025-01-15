# Worker Pools
```go
// ဒီဥပမာမှာ goroutine တွေနဲ့ channel တွေကို သုံးပြီး
// _worker pool_ တစ်ခုကို ဘယ်လို အကောင်အထည်ဖော်မလဲဆိုတာကို ကြည့်ကြပါမယ်။

package main

import (
	"fmt"
	"time"
)

// ဒါက worker ပါ၊ ကျွန်တော်တို့က ဒီ worker ရဲ့ instance အများကြီးကို
// တပြိုင်နက်တည်း run မှာပါ။ ဒီ worker တွေက `jobs` channel ကနေ
// အလုပ်တွေကို လက်ခံပြီး ရလဒ်တွေကို `results` channel ဆီ ပို့ပါလိမ့်မယ်။
// expensive task အလုပ်တစ်ခုကို simulate လုပ်ဖို့ task တစ်ခုအတွက်
// တစ်စက္ကန့် sleep ပါလိမ့်မယ်။
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// ကျွန်တော်တို့ရဲ့ worker pool ကို သုံးဖို့ဆိုရင် သူတို့စီ အလုပ်တွေပို့ပြီး
	// သူတို့ရဲ့ ရလဒ်တွေကို စုစည်းဖို့ လိုပါတယ်။ ဒီအတွက် channel နှစ်ခု လုပ်ပါမယ်။
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// ဒါက worker သုံးခုကို စတင်ပါတယ်၊ ဒါပေမယ့် အစပိုင်းမှာ အလုပ်မရှိသေးတဲ့အတွက်
	// block ဖြစ်နေပါလိမ့်မယ်။
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// ဒီမှာ အလုပ် 5 ခု ပို့ပြီးတော့ အဲဒီ channel ကို `close` လုပ်ပါတယ်။
	// ဒါက ရှိသမျှအလုပ်အားလုံး ပို့ပြီးပြီဆိုတာကို ညွှန်ပြတာပါ။
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// နောက်ဆုံးမှာ အလုပ်အားလုံးရဲ့ ရလဒ်တွေကို စုစည်းပါတယ်။
	// ဒါက worker goroutine တွေ ပြီးဆုံးသွားပြီဆိုတာကိုလည်း သေချာစေပါတယ်။
	// Goroutine အများကြီးကို စောင့်ဆိုင်းဖို့ နောက်ထပ်နည်းလမ်းတစ်ခုကတော့
	// [WaitGroup](waitgroups) ကို သုံးတာပါ။
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
```
```sh
# ကျွန်တော်တို့ run နေတဲ့ပရိုဂရမ်က 
# အလုပ် 5 ခုကို worker တွေက
# executed နေပါတယ်။ 
# စုစုပေါင်း အလုပ်ချိန် 5 စက္ကန့်လောက် 
# ရှိသော်လည်း ပရိုဂရမ်က 
# 2 စက္ကန့်လောက်ပဲ ကြာပါတယ်။ 
# ဘာကြောင့်လဲဆိုတော့ 
# worker 3 ခုက 
# တပြိုင်နက်တည်း အလုပ်လုပ်နေလို့ပါ။
$ time go run worker-pools.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real	0m2.358s
```
