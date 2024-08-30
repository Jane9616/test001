package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		// fmt.Println("开始fibonacci函数for循环")
		select {
		case c <- x:
			// fmt.Println("已完成接收 c <- x")
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main6() {
	c := make(chan int)
	quit := make(chan int)

	// fibonacci若在go gorunting前调用，则函数中的【c <- x】会一直阻塞，导致main函数无法退出
	// fibonacci(c, quit)

	go func() {
		for i := 0; i < 10; i++ {
			// fmt.Println("开始gorunting函数for循环，i = ",i)
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}