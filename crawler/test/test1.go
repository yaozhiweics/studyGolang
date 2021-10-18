package main

import (
	"fmt"
	"time"
)

// 生产者: 生成 factor 整数倍的序列
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		fmt.Println("生产者生产+", i)
		out <- i * factor
	}
}

// 消费者
func Consumer(in <-chan int) {
	for {
		fmt.Println("消费者消费", <-in)
	}

	//for v := range in {
	//	fmt.Println(v)
	//}
}
func main33() {
	ch := make(chan int, 64) // 成果队列

	Producer(1, ch) // 生成 3 的倍数的序列
	//Producer(5, ch) // 生成 5 的倍数的序列
	go Consumer(ch) // 消费 生成的队列

	// 运行一定时间后退出
	time.Sleep(5 * time.Second)
}
