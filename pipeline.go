package petl

import (
	"strings"
	"sync"
	"sync/atomic"
)

type pipeline func(<-chan string) <-chan string

//Spawn :N routines, after each completes runs all whendone functions
func Spawn(N int, fn func(), whendone ...func()) {
	waiting := int32(N)
	for k := 0; k < N; k++ {
		go func() {
			fn()
			if atomic.AddInt32(&waiting, -1) == 0 {
				for _, fn := range whendone {
					fn()
				}
			}
		}()
	}
}

//PipeProcess :
func PipeProcess(initData <-chan string, pipelines ...pipeline) <-chan string {
	data := initData
	for _, fn := range pipelines {
		data = fn(data)
	}
	return data
}

//Extract :
func Extract(strs ...string) <-chan string {
	out := make(chan string, len(strs))
	for _, n := range strs {
		out <- n
	}
	close(out)
	return out
}

//TransformRemoveSpace :
func TransformRemoveSpace(in <-chan string) <-chan string {
	out := make(chan string, len(in))
	for n := range in {
		ret := strings.Replace(n, " ", "", -1)
		out <- ret
	}
	close(out)
	return out
}

func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int, len(in))
	for n := range in {
		out <- n * n
	}
	close(out)
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
