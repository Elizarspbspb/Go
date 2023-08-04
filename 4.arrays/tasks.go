package main

import "fmt"

func main () {
	// Task 1: на вход число N(N≥4), а затем N целых чисел среза. На вывод 4-ый элемент среза.
	fmt.Println("-------- Task 1 --------")
	var N, number int
	fmt.Scanln(&N)
	var slice []int
	for i:=0; i<N; i++ {
		fmt.Scan(&number)
		slice = append(slice, number)
	}
	fmt.Println(slice[3])
}