package main

import "fmt"

func main () {
	// Task 1: на вход число N(N≥4), а затем N целых чисел среза. На вывод 4-ый элемент среза.
	fmt.Println("-------- Task 1 --------")
	/*var N, number int
	fmt.Scanln(&N)
	var slice []int
	for i:=0; i<N; i++ {
		fmt.Scan(&number)
		slice = append(slice, number)
	}
	fmt.Println(slice[3])*/

	// Task 2: На ввод подаются пять целых чисел в массив. найти максимальное число в этом массиве.
	fmt.Println("-------- Task 2 --------")
	array := [5]int{}
	var input, max int
	for i:=0; i < 5; i++{
		fmt.Scan(&input)
		array[i] = input
		if max == 0 {
			max = array[i]
		}
		if max < array[i] {
			max = array[i]
		}
	}
	fmt.Println(max)
	// Task 3: Дан массив из целых чисел. Вывести четный элементы массива
	fmt.Println("-------- Task 3 --------")
	var Num, input_n int
	fmt.Scanln(&Num)
	even_array := []int{}
	for i:=0; i < Num; i++{
		fmt.Scan(&input_n)
		even_array = append(even_array, input_n)
		if i % 2 == 0 {
			fmt.Print(even_array[i], " ")
		}
	}
}