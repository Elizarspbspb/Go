package main

import "fmt"

func main() {
	// Task 1: input and output word
	fmt.Println("-------- Task 1 --------")
    lang := ""
    fmt.Scanln(&lang) // input: color
    fmt.Println(lang) // output: color
	
	// Task 2: 3 numbers , count sum and multiplication on 2 lines
	fmt.Println("-------- Task 2 --------")
	var number_1, number_2, number_3 int
	fmt.Scan(&number_1, &number_2, &number_3)
	fmt.Println(number_1 + number_2 + number_3)
	fmt.Println(number_1 * number_2 * number_3)

	// Task 3: input 1 integer number , output mark telephone if/else
	fmt.Println("-------- Task 3 --------")
	var cost int
	fmt.Scan(&cost)
	if(cost > 1000) {
		fmt.Println("Apple")
	} else if(cost >= 500 && cost <= 1000) {
		fmt.Println("Samsung")
	} else {
		fmt.Println("Nokia")
	}

	// Task 4: input 1 integer number , output factorial (!) until number
	fmt.Println("-------- Task 4 --------")
	var number int
	count, result := 1, 1
	fmt.Scanln(&number)
	for ; count <= number; count++ {
		result *= count
	}
	fmt.Println(result) //

	// Task 5: input 3 integer number with new line, output this 3 number in view text
	fmt.Println("-------- Task 5 --------")
	var number__1, number__2, number__3, text int
	fmt.Scan(&number__1, &number__2, &number__3)
	for i := 0; i < 3; i++ {
		if (i == 0) {
			text = number__1 
		} else if (i == 1) {
			text = number__2 
		} else {
			text = number__3 
		}
		switch text {
			case 0: 
				fmt.Println("Ноль")
			case 1: 
				fmt.Println("Один")
			case 2: 
				fmt.Println("Два")
			case 3: 
				fmt.Println("Три")
			case 4: 
				fmt.Println("Четыре")
			case 5: 
				fmt.Println("Пять")
			case 6: 
				fmt.Println("Шесть")
			case 7: 
				fmt.Println("Семь")
			case 8: 
				fmt.Println("Восемь")
			case 9: 
				fmt.Println("Девять")
			case 10: 
				fmt.Println("Десять")
			default:

		}
	}
}