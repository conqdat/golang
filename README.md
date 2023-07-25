# Learn golang with love

# 7. ****Goroutine****

> A *goroutine* is a lightweight thread of execution.
>
- Explanation

```jsx
package main

import (
	"fmt"
	"sync"
	"time"
)

// main() is also have goroutine, when main goroutine stop, program is also stop.
func main() {
	var wg = sync.WaitGroup{}
	wg.Add(2) // It means Wait Group will wait for 2 goroutine
	// go count("Dat ") // using go keyword to create goroutine

	go func() {
		count("Thao")
		wg.Done()
	}()

	go func() {
		count("DaT")
		wg.Done()
	}()
	wg.Wait() // this is a wait point
	// time.Sleep(time.Second * 5) // Sleep is not good
}

func count(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(name, i)
		time.Sleep(time.Second) // wait for a second
	}
}
```