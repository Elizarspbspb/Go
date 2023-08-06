package main

import "fmt"

func main() {
	var a rune = 's'		// 115
	fmt.Printf("%q", a) 	// s
	fmt.Println("%q", a) 	// %q 115
	fmt.Print("\n%t - ")
    fmt.Printf("%t", a) 	// вывод типа boolean (true или false)
	fmt.Print("\n%b - ")
    fmt.Printf("%b", a) 	// вывод целых чисел в двоичной системе
	fmt.Print("\n%c - ")
    fmt.Printf("%c", a)		// вывод символов, представленных числовым кодом
	fmt.Print("\n%d - ")
    fmt.Printf("%d", a)		// вывод целых чисел в десятичной системе
	fmt.Print("\n%o - ")
    fmt.Printf("%o", a)		// вывод целых чисел в восьмеричной системе
	fmt.Print("\n%q - ")
    fmt.Printf("%q", a)  	// вывод символов в одинарных кавычках
	fmt.Print("\n%x - ")
    fmt.Printf("%x", a)  	// вывод целых чисел в шестнадцатеричной системе, буквенные символы числа имеют нижний регистр a-f
	fmt.Print("\n%X - ")
    fmt.Printf("%X", a)		// вывод целых чисел в шестнадцатеричной системе, буквенные символы числа имеют верхний регистр A-F
	fmt.Print("\n%U - ")
    fmt.Printf("%U", a)		// вывод символов в формате кодов Unicode, U+1234
	fmt.Print("\n%e - ")
    fmt.Printf("%e", a)		// вывод чисел с плавающей точкой в экспоненциальном представлении, -1.234456e+78
	fmt.Print("\n%E - ")
    fmt.Printf("%E", a)		// аналог %e но в верхнем регистре, -1.234456E+78
	fmt.Print("\n%f - ")
    fmt.Printf("%f", a)		// вывод чисел с плавающей точкой, например, 123.456
	fmt.Print("\n%F - ")
    fmt.Printf("%F", a)		// то же самое, что и %f
	fmt.Print("\n%g - ")
    fmt.Printf("%g", a)		// %g   %e для огромных экспонент, %f в противном случае
	fmt.Print("\n%e - ")
	fmt.Printf("%e", a)		// %g   %e для огромных экспонент, %f в противном случае
	fmt.Print("\n%G - ")
    fmt.Printf("%G", a) 	// %G   %E для огромных экспонент, %F в противном случае
	fmt.Print("\n%E - ")
	fmt.Printf("%E", a)		// %G   %E для огромных экспонент, %F в противном случае
    fmt.Print("\n%s - ")
	fmt.Printf("%s", a)		// вывод строки
	fmt.Print("\n%p - ")
    fmt.Printf("%p", a)		// вывод значения указателя - адреса в шестнадцатеричном представлении
	fmt.Print("\n%T - ")
    fmt.Printf("%T", a) 	// вывод типа переменной */
	fmt.Println()

	var number = 1234.6789123
	fmt.Printf("%f\n", number)		// 1234.678912
	fmt.Printf("%9.f\n", number)	// _____1235 9-позиций.
	fmt.Printf("%9.2f\n", number)	// __1234.68 6.2=9

	var num float64 = 100.123456
	fmt.Printf("это число %f типа %T\n", num, num)	// это число 100.123456 типа float64
	var a1 byte = 's'
	var a2 int = 1234
	fmt.Printf("%q %b\n", a1, a2)	// 's' 10011010010

	fmt.Println(`Sammy says, "Hello!"`)	// Sammy says, "Hello!"
	fmt.Println("Say \"hello\" to Go!")	// Say "hello" to Go!
	fmt.Println("Sammy" + "Shark")		// SammyShark
	fmt.Println("This string\nspans multiple\nlines.")
	fmt.Println(`Sammy says,\"The balloon\'s color is red.\"`) 	// Sammy says,\"The balloon\'s color is red.\" 

	// Task 1: На вход подается число типа float64. Вывести преобразованное число: число возводится в квадрат, 
	// затем обрезается дробная часть так что остается 4 знака после запятой. Но если число равно или меньше нуля
	// - выводить: "число R не подходит", где R - исходное число с 2 цифрами после запятой и с 2 по ширине. 
	// А если число больше чем 10 000 - выводить исходное число в экспоненциальном представлении (знак экспоненты в нижнем регистре).
	fmt.Println("-------- Task 1 --------")
	var input float64
	fmt.Scanln(&input)
	if input > 10000 {
		fmt.Printf("%e", input)	
	} else if input <= 0 {
		fmt.Printf("число %2.2f не подходит", input)
		// "число не подходит" - литерал
		// число не подходит - значение
	} else {
		fmt.Printf("%.4f\n", input*input)
	}
	// ///////////////////////////////
	result := fmt.Sprintf("%.2f", input)// ничего не выводит
	fmt.Printf("%q", result) 			// вывод: "100.12"	
	// result будет типа string
}