package basic_core

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func InitContext() {
	goWork1()
}

var wg sync.WaitGroup

func goWork1() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker1(ctx)
	time.Sleep(time.Second * 3)
	cancel()
	wg.Wait()
	fmt.Println("over")

}

func worker1(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker1")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	wg.Done()
}

func goWorker() {
	var exitChan = make(chan struct{})
	wg.Add(1)
	go worker(exitChan)
	time.Sleep(3 * time.Second)
	exitChan <- struct{}{}
	close(exitChan)
	wg.Wait()
	fmt.Println("over")
}

func worker(exitChan chan struct{}) {
LOOP:
	for {
		fmt.Println("init worker")
		time.Sleep(time.Second)
		select {
		case <-exitChan:
			break LOOP
		default:
		}
	}
	wg.Done()
}
