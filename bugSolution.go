```go
package main

import (


    "fmt"
    "sync"
)

func main() {
    var m sync.Mutex
    var wg sync.WaitGroup
    ch := make(chan int)

    wg.Add(1)
    go func() {
        defer wg.Done()
        m.Lock()
        fmt.Println("goroutine1 acquired lock")
        m.Unlock()
        close(ch)
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()
        <-ch
        fmt.Println("goroutine2 received from channel")
        m.Lock()
        fmt.Println("goroutine2 acquired lock")
        m.Unlock()
    }()

    wg.Wait()
}
```