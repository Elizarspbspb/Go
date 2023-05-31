package main

import "fmt"

func main() {

	var number_1 int = 30
	var number_2 int = 15
    
	fmt.Println(number_1 == number_2) // output false
	fmt.Println(number_1 != number_2) // output true
	fmt.Println(number_1 > number_2)  // output true
	fmt.Println(number_1 >= number_2) // output true
	fmt.Println(number_1 < number_2)  // output false
	fmt.Println(number_1 <= number_2) // output false
    fmt.Println("------------")
	fmt.Println(number_1 != number_2 && number_1 >= number_2)  	// output true
	fmt.Println(number_1 == number_2 || number_1 > number_2)   	// output true
	fmt.Println(!(number_1 < number_2))          				// output true
	fmt.Println("------------")	
	fmt.Println(!(8 >= 5) && 6 == 6) 		
	fmt.Println(!(12 > 4 || 8 != 5)) 		
}