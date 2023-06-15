package main

import "fmt"

func main () {
	// ------------- 1 Three-component loop -----------------
	sum_1 := 0
	for i := 1; i < 5; i++ {
		sum_1 += i
	}
	fmt.Println(sum_1) // 10 (1+2+3+4)

	// ------------- 2 While loop -----------------
	number := 1
	for number < 5 {
		number *= 2
	}
	fmt.Println(number) // 8 (1*2*2*2)

	// ------------- 3 Infinite loop -----------------
	/*sum_2 := 0
	for {
		sum_2++ // repeated forever
	}
	fmt.Println(sum_2) // never reached */

	// ------------- 4 For-each range loop -----------------
	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	// ------------- 5 Exit a loop -----------------
	sum := 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 { // skip odd numbers
			continue
		}
		sum += i
	}
	fmt.Println(sum) // 6 (2+4)		
}