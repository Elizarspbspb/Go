package main

import (
    "fmt"
    "time"
)

func out(from, to int, ch chan bool) {
    for i:=from; i<=to; i++ {
        time.Sleep(50 * time.Millisecond)
        fmt.Print(i)
    }
	fmt.Println()
    ch <- true
}

func evenSum(from, to int, ch chan int) {
	result := 0
	for i:=from; i<=to; i++ {
		if i%2 == 0 {
			result += i
		}    
	}
	ch <- result
}
func squareSum(from, to int, ch chan int) {
	result := 0
	for i:=from; i<=to; i++ {
	  	if i%2 == 0 {
			result += i*i
	  	}    
	}
	ch <- result
} 

func main() {
  	ch := make(chan bool)	// create channel

  	go out(0, 4, ch)		// Parallel start 3 function (2 - out, 1 - main)
  	go out(6, 9, ch)		// Parallel start, but main finished early than outs

  	<-ch		// waiting data from the channel
  	<-ch		// waiting data from the channel
	close(ch)	// close channel

  	evenCh := make(chan int)
  	sqCh := make(chan int)

  	go evenSum(0, 100, evenCh)
  	go squareSum(0, 100, sqCh)

  	fmt.Println(<-evenCh + <-sqCh)
}