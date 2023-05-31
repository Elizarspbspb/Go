package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	fmt.Println(number) 	// output number
	fmt.Println("\n------------")	
	var number_1, number_2, number_3, number_4 int
	fmt.Scan(&number_1, &number_2)
	fmt.Print(number_1 + number_2)
	fmt.Println("\n------------")	
	number_1 = 0
	number_2 = 0
    fmt.Scanln(&number_1, &number_2, &number_3, &number_4) // only line
    fmt.Print(number_1, number_2, number_3, number_4)
}