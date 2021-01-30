package main

import (
	"fmt"
	"sync"
)

type concurrentSingle struct {
	i int
}

func NewSingleton() func() *concurrentSingle {
	once := &sync.Once{}
	var p *concurrentSingle
	return func() *concurrentSingle {
		once.Do(func() {
			p = new(concurrentSingle)
		})
		p.i++
		return p
	}
}

func main() {
	wg := &sync.WaitGroup{}
	createOnce := NewSingleton()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(createOnce().i)
			wg.Done()
		}()
	}

	wg.Wait()
}
