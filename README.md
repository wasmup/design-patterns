
## Concurrent Design Patterns in the Go (Golang)


### Concurrent Singleton Design Pattern in the Go (Golang)

Try it on [The Go Playground](https://play.golang.org/p/UGZ4k6Gli39)
```go
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

```
Output:
```sh
1
2
4
5
3
6
7
9
8
10
```

----

## Singleton2
Try it on [The Go Playground](https://play.golang.org/p/lo0Xut9rP4V)
```go
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
```