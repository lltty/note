package main

import (
	"fmt"
	"sync"
	"time"
)

var msgChannel chan int
var msgCount int
var interval time.Duration

func init() {
	msgCount = 18
	msgChannel = make(chan int, msgCount)
	interval = 1500 * time.Millisecond
}

func receive(queue <-chan int, wg *sync.WaitGroup) {
	fmt.Println("receive msg ...")
	expire := time.After(30 * time.Second)
	defer wg.Done()
LOOP:
	for {
		select {
		case msg, ok := <-queue:
			if !ok {
				fmt.Println("close chan")
				break LOOP
			}
			fmt.Printf("receivce msg %d\n", msg)
		case <-expire:
			fmt.Println("timeout!!!")
			break LOOP
		}
	}
}

//这个函数可以起到热替换的效果
func getChan() chan int {
	return msgChannel
}

func receivePerf(wg *sync.WaitGroup) {
	fmt.Println("receive msg ...")
	defer wg.Done()
LOOP:
	for {
		expire := time.After(10 * time.Millisecond)
		select {
		/*
		 * 注意这种热切换也是有问题的，这必须要在代码02等待的时间比代码01等待的时间久的情况
		 * 因为当msgChannel在代码03的地方被重置的之后，getChan()函数的执行可能在重置之前，那么获取到的依然是旧的
		 * channel值，这时候就会导致阻塞，然后一直获取不到值
		 */
		case msg, ok := <-getChan():
			if !ok {
				fmt.Println("close chan")
				break LOOP
			}
			fmt.Printf("receivce msg %d\n", msg)
			//代码02
			time.Sleep(1500 * time.Millisecond)
		case <-expire:
			fmt.Println("timeout!!!")
			//break LOOP
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < msgCount; i++ {
			//没三次对channel做一次重置
			if i > 0 && i%3 == 0 {
				fmt.Println("reset chan")
				//代码03
				msgChannel = make(chan int, msgCount)
			}

			fmt.Printf("send msg %d \n", i)
			msgChannel <- i
			//每次操作等待1.5秒
			//代码01
			time.Sleep(interval)
		}
		//所有操作完成之后关闭通道
		fmt.Println("close channel")
		close(msgChannel)
		wg.Done()
	}()

	wg.Add(1)
	/*
	 * msgChannel在上面被替换，但是receive函数中的接收操作缺仍然是之前的那个通道，
	 * 这样就使本该操作同一通道值的两个操作不匹配了。
	 * 所以receive只能接收到channel重置之前的0, 1, 2, 3
	 */
	//go receive(msgChannel, &wg)

	//热替换
	go receivePerf(&wg)
	wg.Wait()
}
