package main

import (
		"fmt"
	)

func main() {
	//var wg *sync.WaitGroup
	//wg = new(sync.WaitGroup)
	//wg.Add(10)
	ch := make(chan int,0)
	for i:=0 ;i<10;i++ {
		go func(ii int) {
				//time.Sleep(time.Second*2)
				fmt.Println("ii:",ii)
				ch <- ii
				//wg.Done()
			}(i)
		}

		for {
			select {
				case x:=<-ch:
					fmt.Println("receive",x)
				default:
			}
		}
	//<- ch // 阻塞main goroutine, 信道c被锁
}
//func main() {
//	c := make(chan int)
//
//	go func() {
//		for i:=0 ;i<10;i++ {
//			time.Sleep(time.Second*2)
//			c <- i
//		}
//	}()
//
//	for {
//		select {
//			case x:=<-c:
//				fmt.Println("axxx",x)
//			default:
//		}
//	}
//}