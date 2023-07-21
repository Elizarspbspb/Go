package main

import ("fmt"
		"math"
)

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

	// Task 6: По данному трехзначному числу определить, все ли его цифры различны
	fmt.Println("-------- Task 6 --------")
	var number___1 string
	fmt.Scan(&number___1)
	if ((number___1[0] == number___1[1]) || (number___1[1] == number___1[2]) || (number___1[0] == number___1[2])) {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}

	// Task 7: Вывести первую цифру числа
	fmt.Println("-------- Task 7 --------")
	// first variant
	var input_number_str string
	fmt.Scan(&input_number_str)
	fmt.Println(input_number_str[0]-48)
	fmt.Println(string(input_number_str[0]))
	// second variant with recursion
	var input_number_int int
	fmt.Scan(&input_number_int)
	fmt.Println(getFirstNumber(input_number_int))
	// first variant with if...else
	if input_number_int <=9 {
        fmt.Println(input_number_int)
    } else if input_number_int >=10 && input_number_int <=99 {
        fmt.Println(input_number_int/10)
    } else if input_number_int >=100 && input_number_int <=999 {
        fmt.Println(input_number_int/100)
    } else if input_number_int >=1000 && input_number_int <=9999 {    
        fmt.Println(input_number_int/1000)
    } else {fmt.Println(input_number_int/10000)}   
	// fourth variant with math 
	var input_number_float float64
	fmt.Scan(&input_number_float)
	fmt.Println(int(input_number_float / math.Pow10(int(math.Log10(input_number_float)))))

	// Task 8: В шестизначном номере сумма первых трёх цифр совпадает с суммой трёх последних Да или Нет
	fmt.Println("-------- Task 8 --------")
	// first variant
	var happy_ticket string
	fmt.Scan(&happy_ticket)
	if((happy_ticket[0] + happy_ticket[1] + happy_ticket[2]) == (happy_ticket[3] + happy_ticket[4] + happy_ticket[5])){
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	// second variant
	var happy_ticket_int  int
	fmt.Scan(&happy_ticket_int)
	if(getSumNumber(happy_ticket_int/1000) == getSumNumber(happy_ticket_int%1000)){
		fmt.Println("YES", getSumNumber(happy_ticket_int/1000), "=", getSumNumber(happy_ticket_int%1000))
	} else {
		fmt.Println("NO")
	}

	// Task 9: Високосный год Да или Нет. Кратен 400, или кратен 4 но не кратен 100
	fmt.Println("-------- Task 9 --------")
	var year int
	fmt.Scan(&year)
	if(((year%400)==0) || ((year%4)==0 && (year%100)!=0)){
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

// Task 7: Вывести первую цифру числа (РЕКУРСИЯ)
func getFirstNumber(a int) int {
	if a < 10 {
		return a
	} else {
	 	return getFirstNumber(a / 10)
	}
}

// Task 8: Вывести сумму 3 цифр в числе
func getSumNumber(a int) int {
	var ost int
	ost = a % 10
	fmt.Println("ost =", ost)
	if a < 10 {
		return a
	} else {
	 	return (ost + getSumNumber(a/10))
	}
}