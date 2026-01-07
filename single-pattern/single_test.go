package single_pattern

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetConfig(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("%p\n", GetConfig())
		}()
	}
	wg.Wait()
}
