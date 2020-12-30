package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	value int64
}

type result struct {
	job *job
	sum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

func generate(j chan<- *job) {
	defer wg.Done()
	//循环生成int64随机数，发送到jobchan
	for {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		j <- newJob
		time.Sleep(time.Second)
	}
}

func _sum(j <-chan *job, resultChan chan<- *result) {
	defer wg.Done()
	for {
		job := <-j
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job: job,
			sum: sum,
		}
		resultChan <- newResult
	}

}

func main() {
	wg.Add(1)
	go generate(jobChan)

	//开启24个gorountine
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go _sum(jobChan, resultChan)
	}

	for res := range resultChan {
		fmt.Printf("value:%d,sum:%d\n", res.job.value, res.sum)
	}
	wg.Wait()
}
