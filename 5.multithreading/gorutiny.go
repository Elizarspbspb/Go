package main

import (
    "fmt"
    "time"
)

func out(from, to int) {
  	for i:=from; i<=to; i++ {
    	time.Sleep(50 * time.Millisecond)
    	fmt.Print(i)
  	}
  	fmt.Println()
}

func main() {
  	out(0, 5)   		// Sequential start
  	out(6, 10)		// Последовательный вывод
  	// 012345
	// 678910
  	go out(0, 5)		// Parallel start 3 function (2 - out, 1 - main)
  	go out(6, 10)		// Parallel start, but main finished early than outs
  	time.Sleep(500 * time.Millisecond)
}