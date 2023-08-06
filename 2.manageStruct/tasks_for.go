package main

import ("fmt"
)

func main() {

	// Task 1: Ввести 2 числа A и B - вывести сумму всех чисел от A до B включительно.
	fmt.Println("-------- Task 1 --------")
	var number_1 int
    var number_2 int
    var count int
    fmt.Scanln(&number_1, &number_2)	// 1 5
    for i:=number_1; i<=number_2; i++{
        count += i
    }
    fmt.Println(count)	// 15

	// Task 2: В последовательности чисел найти сумму 2-значных и кратных 8. 
	fmt.Println("-------- Task 2 --------")
    var input_count, input_num, until, sum int
    fmt.Scanln(&input_count)										// 5
	for fmt.Scan(&input_num); until < input_count; fmt.Scan(&input_num) {	// 38 24 800 8 16
		if input_num % 8 == 0 && input_num < 100 && input_num > 9 {
			sum += input_num
		}
		until++
		if until == 5 {
			break
		}
	}
	/////////////
	/*for fmt.Scan(&n); n > 0; n-- {
		var x int
		if fmt.Scan(&x); x > 9 && x < 100 && x%8 == 0 {
			sum += x
		}
	}*/
	// ! ! ! fmt.Println(sum)	// 40

	// Task 3: В последовательности определить количество элементов, которые равны ее наибольшему элементу.
	fmt.Println("-------- Task 3 --------")
	var input, count_max, max int	
	for fmt.Scan(&input); input != 0; fmt.Scan(&input) {	// 1 2 3 4 4 4 3 0
		if input > max {
			max = input
			count_max = 1
		} else if input == max {
			count_max++
		}
	}
	fmt.Println(count_max)	// 3

	// Task 4: Вводится 3 натуральных числа n, c, d. Вывести первое число от 1 до n кратное c, но НЕ кратное d.
	fmt.Println("-------- Task 4 --------")
	var n, c, d int
    fmt.Scanln(&n)
    fmt.Scanln(&c)
    fmt.Scanln(&d)
    i := 1
    for i<=n { 
        if (i%c == 0) && (i%d != 0) {
            fmt.Println(i)
            break
        }
        i++
    }

	// Task 5: Считывать числа. Если < 10 пропустить, если > 100 прервать цикл.
	fmt.Println("-------- Task 5 --------")
	var input_s int	
	for {			
		fmt.Scanln(&input_s)
		if input_s < 10 {
			continue
		} else if input_s > 100 {
			break
		}
		fmt.Println(input_s)
	}

	// Task 6: Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
	fmt.Println("-------- Task 6 --------")
	var num_1, num_2 string
	fmt.Scanln(&num_1, &num_2)
	for i := 0; i < len(num_1); i++ {
		for j := 0; j < len(num_2); j++ {
			if num_1[i] == num_2[j] {
				fmt.Println(num_1[i] - 48)
				fmt.Println(" ")
			}
		}
	}
}