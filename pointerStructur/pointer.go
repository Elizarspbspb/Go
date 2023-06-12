package main

import "fmt"

func change(num int) { // copy num
    num = 22
}

func changePointer(pointer *int) { // number origin
    *pointer = 22
}

func changeStrings(str_1 *string, str_2 *string) {
	str := ""
	str = *str_1
	*str_1 = *str_2
	*str_2 = str 
}

func main() {
	num := 26
	var point *int // Pointers are special variables that have the addresses of values in memory
	point = &num
	//point := &num
	fmt.Println(*point) // output 26
	*point = 15
	fmt.Println(*point) // output 15
	fmt.Println(num)  	// output 15
	fmt.Println(point) 	// output 0xc0000ba000
	fmt.Println(&num)  	// output 0xc0000ba000
	fmt.Println(&point) // output 0xc0000b4018

	change(num)
    fmt.Println(num) 	// output 15
	changePointer(&num)
    fmt.Println(num) 	// output 22

	str_f := "First"
	str_s := "Second"
	fmt.Println(str_f, str_s) 	// First Second
	//changeStrings(&"First", &"Second") // cannot take address of "First"
	changeStrings(&str_f, &str_s)
	fmt.Println(str_f, str_s) 	// Second First
}