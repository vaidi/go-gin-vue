package go_runtime

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func counter(out chan<- string) {
	for i := 0; i < 10; i++ {
		out <- fmt.Sprintf("%d", i)
	}
	close(out)
}
func square(out chan<- string, in <-chan string) {
	for i := range in {
		out <- "haha:" + i
	}
	close(out)
}
func printer(in <-chan string) {
	for s := range in {
		println(s)
	}
}

func SingleChannel() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go square(ch2, ch1)
	go counter(ch1)
	printer(ch2)
}

func GoChannelGrace() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
			fmt.Println("i:" + strconv.Itoa(i))
		}
		close(ch2)
	}()
	//从渠道中一种特有的遍历方式
	for i := range ch2 {
		fmt.Println("得到的i:" + strconv.Itoa(i))
	}

}

func GoChannelCaChe() {
	ch := make(chan string, 1)
	ch <- "my love"
	a := <-ch
	ch <- "my love"
	fmt.Printf("%v xxh", a)

}

func GoChannel() {
	var ch chan int
	fmt.Println(ch)
	ch = make(chan int)
	fmt.Println(ch)

	go func() {
		x := <-ch
		fmt.Println("拿到渠道" + strconv.Itoa(x))
	}()
	ch <- 10
	runtime.Gosched()
	defer func() {
		close(ch)
	}()

}

func GoMaxProc() {
	println("cpu:" + strconv.Itoa(runtime.NumCPU()))
	runtime.GOMAXPROCS(2)
	go a()
	go b()

}
func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("a:", i)
	}
}
func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("b:", i)
	}
}

func GoExit() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			//退出 就是整个方法都不在执行
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("D.defer")
		}()
		fmt.Println("E.defer")
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("H.defer")
}

func GoRunTime() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			println("s:" + s + ":" + strconv.Itoa(i))
		}
	}("hello world")
	for i := 0; i < 2; i++ {
		//让出时间片
		runtime.Gosched()
		fmt.Println("hello goSched")
	}
}
