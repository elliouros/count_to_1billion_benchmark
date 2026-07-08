package main

import (
	"fmt"
	"time"
)

func count_to_billion() uint64 {
	var i uint64
	for i = 0; i < 1000000000; i++ {
	}
	return i
}

func main() {
	start := time.Now()
	count := count_to_billion()
	duration := time.Since(start).Seconds()
	fmt.Printf("Looping %d times in go took: %f\n", count, duration)
}
