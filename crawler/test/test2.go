package main

import (
	"fmt"
	"time"
)

// 生产者
func producer(channel chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("老怪创造了小怪", i)
		time.Sleep(time.Second)
		channel <- i
	}

	// 关闭通道
	close(channel)
}

// 消费者
func consumer(name string, channel chan int, done chan bool) {
	for {
		val, ok := <-channel
		if ok {
			fmt.Println(name, "消灭了怪物", val)
		} else {
			fmt.Println("消灭了所有怪物！！")
			break
		}
	}

	// ok
	done <- true
}

// 消费者
func consumer2(name string, channel chan int) {
	for {
		val, ok := <-channel
		if ok {
			fmt.Println(name, "消灭了怪物", val)
		} else {
			fmt.Println("消灭了所有怪物！！")
			break
		}
	}

	// ok
	// done <- true
}

func maind() {
	c1 := make(chan int)
	//done := make(chan bool)

	go producer(c1)
	go consumer2("陈大雷", c1)
	go consumer2("迪丽热巴", c1)

	//<- done
}
