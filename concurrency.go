package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 2; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(n []int, c chan int) {
	sum := 0
	for _, value := range n {
		sum += value
	}
	c <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
			fmt.Println("Produce:", x)
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}

// SafeCounter的并发使用是安全的
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

//Inc 增加给定key的计数器的值
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	//Lock之后同一时刻只有一个goroutine能访问c.v
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	defer fmt.Println("Bye")
	go say("world")
	say("Hello")

	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	c := make(chan int)
	go sum(s[:4], c)
	go sum(s[4:], c)
	x := <-c
	y := <-c
	fmt.Println(x, y, x+y)

	ch := make(chan int, 8)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	ch <- 6
	ch <- 7
	ch <- 8
	fmt.Print(<-ch)
	fmt.Print(<-ch)
	fmt.Print(<-ch)
	fmt.Print(<-ch)
	fmt.Print(<-ch)
	fmt.Print(<-ch)
	fmt.Print(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch)    //fatal error: all goroutines are asleep - deadlock!

	ch2 := make(chan int)
	go fibonacci(10, ch2)
	for i := range ch2 {
		fmt.Println(i)
	}

	ch3 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Rec:", i, <-ch3)

		}
		quit <- 0
	}()
	fibonacci2(ch3, quit)

	// sync.Mutex
	cc := SafeCounter{v: make(map[string]int)}
	//fmt.Println("cc", cc)  //call of fmt.Println copies lock value: play.SafeCounter contains sync.Mutex
	for i := 0; i < 1000; i++ {
		go cc.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(cc.Value("somekey"), time.Second)

	// select
	tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(5000 * time.Millisecond)
	for {
		select {
		case counter := <-tick:
			fmt.Print("counter.", counter.Second())
		case <-boom:
			fmt.Println("Boom!")
			return
		default:
			fmt.Print(".")
			time.Sleep(500 * time.Millisecond)

		}
	}

}
