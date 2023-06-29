package main

import "fmt"

func main() {
	var number1 int = 10
	number2 := 15
	var result int
      
	result = number1 + number2 	// addition
	fmt.Println(result)  		// output 25
	result = number1 - number2 	// subtraction
	fmt.Println(result)  		// output -5
	result = number1 * number2 	// multiplication
	fmt.Println(result)  		// output 150
	result = number2 / number1 	// division
	fmt.Println(result)  		// output 1
	result = number2 % number1 	// mod 
	fmt.Println(result)  		// output 5

	var x_f float32 = 13/4.0
	fmt.Println(x_f)     // output: 3.25
	var x_1 = 9 / 2.0
	fmt.Println(x_1)	// output: 4.5
	x_2 := 9 / 2.0
	fmt.Println(x_2)	// output: 4.5

	number2 += number1
	fmt.Println(number2)	// output: 25

	var incr int = 2
	incr++            
	fmt.Println(incr) 	// output 3
	incr--				// --incr - ERROR
	fmt.Println(incr) 	// output 2

	variable_1 := "Привет, "
	variable_2 := "Степик!"
	fmt.Println(variable_1 + variable_2) 	// Concatenation: Привет, Степик!
}