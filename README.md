
## Concurrent Design Patterns in the Go (Golang)


### Concurrent Singleton Design Pattern in the Go (Golang)

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
