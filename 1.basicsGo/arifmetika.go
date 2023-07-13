package main

import "fmt"

func main() {
	//var 1_number int = 10 // ERROR. First symbol only _ or letter (_ или буква)
	var number1 int = 10
	_number2 := 15
	var result int

	// ways to declare variables
	var a,b,c int
	var a int
	a := 15
	var a int = 10
	var a = 12 
    //

	result = number1 + _number2 	// addition
	fmt.Println(result)  			// output 25
	result = number1 - _number2 	// subtraction
	fmt.Println(result)  			// output -5
	result = number1 * _number2 	// multiplication
	fmt.Println(result)  			// output 150
	result = _number2 / number1 	// division
	fmt.Println(result)  			// output 1
	result = _number2 % number1 	// mod 
	fmt.Println(result)  			// output 5

	var x_f float32 = 13/4.0
	fmt.Println(x_f)     	// output: 3.25
	var x_1 = 9 / 2.0
	fmt.Println(x_1)		// output: 4.5
	x_2 := 9 / 2.0
	fmt.Println(x_2)		// output: 4.5

	_number2 += number1
	fmt.Println(_number2)	// output: 25

	var incr int = 2
	incr++            
	fmt.Println(incr) 	// output 3
	incr--				// --incr - ERROR
	fmt.Println(incr) 	// output 2

	variable_1 := "Привет, "
	variable_2 := "Степик!"
	fmt.Println(variable_1 + variable_2) 	// Concatenation: Привет, Степик!

	var symbol int32 = 'c'						// хранение кода символа, а не сам символ
	fmt.Println(symbol, " = ", string(symbol))	// вывод символа соответ. коду | 99 = c 

	fmt.Print("Ivan", 27) 	// Ivan27  	; without probels - so string
	fmt.Println("Ivan", 27) // Ivan 27 	; probel - so Println
	fmt.Print(33, 27) 		// 33 27	; probel - so not string

	// Найти число десятков в числе. 2010 ввели - вывести 1
	var a, b int
	fmt.Scan(&a)
	b = a % 10
	a = a % 100
	fmt.Println((a - b)/10)
	//

	/* Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас 
	целых часов h и целых минут m. Входные данные d (0 < d < 360).
	Input: 	90
	Output: It is 3 hours 0 minutes. 	*/
	var a, h, m int
	fmt.Scan(&a)
	a *= 2
	m = a % 60
	h = (a - m) / 60
	fmt.Println("It is", h, "hours", m, "minutes.")

}