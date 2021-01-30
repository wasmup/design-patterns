package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	i int
}

// Inc adds one to the i and returns the result (just an example).
func (p *singleton) Inc() int {
	p.i++
	return p.i
}

type concurrentSingle struct {
	sync.Once
	single *singleton
}

func (p *concurrentSingle) createOnce() *singleton {
	p.Do(func() {
		p.single = new(singleton)
	})
	return p.single
}

func main() {
	wg := &sync.WaitGroup{}
	p := &concurrentSingle{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(p.createOnce().Inc())
			wg.Done()
		}()
	}

	wg.Wait()
}
