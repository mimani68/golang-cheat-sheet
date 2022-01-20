package goroutine

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("salam")
	}()
	time.Sleep(2 * time.Second)
}
