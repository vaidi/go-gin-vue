package go_runtime

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Job struct {
	Id int
	//随机数
	RandNum int
}
type Result struct {
	job *Job
	sum int
}

func GoTimer() {

	timer := time.NewTicker(2 * time.Second)
	go func() {
		for {
			<-timer.C
			fmt.Println("每2s执行一次")
		}

	}()

}

func GoJobRoutine() {
	jobChan := make(chan *Job, 128)
	resultChan := make(chan *Result, 128)
	createPool(64, jobChan, resultChan)
	go func(resultChan chan *Result) {
		for result := range resultChan {
			fmt.Println("job id:%v randum:%v result:%d\n", result.job.Id, result.job.RandNum, result.sum)
		}
	}(resultChan)
	var id int
	for {
		id++
		r_num := rand.Int()
		job := &Job{id, r_num}
		jobChan <- job
	}
}
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			for job := range jobChan {
				r_num := job.RandNum
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				r := &Result{job, sum}
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}

func GoRoutine() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()

}
func hello(i int) {
	defer wg.Done()
	fmt.Println("hello goroutine", i)
}
