package main

import (
    "fmt"
	"time"
)

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
	evenCh := make(chan int)
	sqCh := make(chan int)

	go evenSum(0, 100, evenCh)
	go squareSum(0, 100, sqCh)		// will be first finish

	select {						// select waits data from channel 
  	case first := <- evenCh:		// and only one of the cases will be executed
     	fmt.Println(first)			// the one that received the data first.
		fmt.Println("First case")
  	case second := <- sqCh:
     	fmt.Println(second)
		fmt.Println("Second case")
	}

	for {
		select {						
		case <- evenCh:					
			fmt.Println("First case")	
		case <- sqCh:
			fmt.Println("Second case")
		default:
            fmt.Println("Nothing is available")
            time.Sleep(50 * time.Millisecond)
		}
	}
	// 171700
	// Second case
	// Nothing is available
	// First case
	// Nothing is available
	// Nothing is available
	// ...
}
