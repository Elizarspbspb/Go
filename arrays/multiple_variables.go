package main

import "fmt"

func su_m(nums ...int) { 	// counting summ elements
	total := 0
	for _, variable := range nums {
	  total += variable
	}
	fmt.Println(total)
  }

func main () {
	fmt.Println("a", "b", "c") // a b c

	su_m(2, 4, 6)          //  12
	su_m(42, 8)            //  50
	su_m(1, 2, 3, 4, 5, 6) //  21

	slice := []int{42, 33, 22, 11}
	su_m(slice...)	// 108
}