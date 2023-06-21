package main

import "fmt"

func main () {
	array := [5]int{-1, 2, 4, 6, 8}

	var slice_1 []int = array[1:3]
	fmt.Println("size = ", len(slice_1), ";") // size =  2 ;
	for i := 0; i < len(slice_1); i++ {	
		fmt.Print(slice_1[i], ", ")		// 2, 4,
	}
	fmt.Print("\n")

	var slice_2 []int = array[:3]
	for i, s := range slice_2 {
		fmt.Print(i, ", ")	// 0, 1, 2
		fmt.Print(s, "; ")	// -1, 2, 4
	}
	fmt.Print("\n")

	slice_2[0] = 100 	// slice change start array
	for i := 0; i < len(array); i++ {	
		fmt.Print(array[i], ", ")		// 100, 2, 4, 6, 8
	}
	fmt.Print("\n")

	// dynamic array
	dyn_array_five := make([]int, 5)
	dyn_array_five = append(dyn_array_five, 8, 16)
	fmt.Println(dyn_array_five) // выведет [0 0 0 0 0 8 16]

	// dynamic array 0
	dyn_array := make([]int, 0)
	dyn_array = append(dyn_array, 8, 16)
	fmt.Println(dyn_array) 		// выведет [8 16]

	for _, s := range dyn_array_five {
		fmt.Print(s, "; ")	// 0; 0; 0; 0; 0; 8; 16
	}
	fmt.Print("\n")

	word := "hello"
	for _, w := range word {
  		fmt.Printf("%c", w) 	// hello
	}
	fmt.Print("\n")
}