/**** Working With Wait Groups ****/

package main
import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("sheep")
		wg.Done()
	}()

	wg.Wait()
}

func count(animal string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, animal)
		time.Sleep(time.Millisecond * 500)
	}
}