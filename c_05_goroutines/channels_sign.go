package main

import (
	"fmt"
	"time"
)

// 信号量模式
func compute(ch chan int){
ch <- 1// when it completes, signal on the channel.
}

func main(){
ch := make(chan int)     // allocate a channel.
go compute(ch)        // stat something in a goroutines
result := <- ch
fmt.Println(result)

    // 匿名函数
	ch1 := make(chan int)
	go func(){
		// doSomething
		ch1 <- 2 // Send a signal; value does not matter
	}()
	fmt.Println(<- ch1)

	// 实现并行的 for 循环
	data := [3]float64{1,2,3}
	for i, v := range data {
		go func (i int, v float64) {
			fmt.Println(i,v)
		} (i, v)
	}
	time.Sleep(1e9)

}

// 多信号量
//done := make(chan bool)
//// doSort is a lambda function, so a closure which knows the channel done:
//doSort := func(s []int){
//    sort(s)
//    done <- true
//}
//i := pivot(s)
//go doSort(s[:i])
//go doSort(s[i:])
//<-done
//<-done
//