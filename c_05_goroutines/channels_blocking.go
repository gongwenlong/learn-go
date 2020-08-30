package main

import (
	"fmt"
	"time"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

//默认情况下，通信是同步且无缓冲的：在有接受者接收数据之前，发送不会结束。
// 可以想象一个无缓冲的通道在没有空间来保存数据的时候：
// 必须要一个接收者准备好接收通道的数据然后发送者可以直接把数据发送给接收者。
// 所以通道的发送 / 接收操作在对方准备好之前是阻塞的：
//
//————————————————
func main() {
	//out := make(chan int, 1)
	out := make(chan int)
	out <- 2
	fmt.Print("qqq")
	go f1(out)
	time.Sleep(1e9)
}