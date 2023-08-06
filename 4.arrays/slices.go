package main

import "fmt"

// срез — указатель на массив
// длина (len) — это количество элементов среза;
// емкость (cap) — количество элементов между началом среза и концом базового массива.

func fnArray(arr [5]int) {
	arr[1] = 15	// copy array
}

func fnSlice(arr []int) {
	arr[1] = 15	// link on array
}

func main () {
	array := [5]int{-1, 2, 4, 6, 8}

	var slice_1 []int = array[1:3]
	fmt.Println("size = ", len(slice_1), ";") // size =  2 ;
	for i := 0; i < len(slice_1); i++ {	
		fmt.Print(slice_1[i], ", ")		// 2, 4,
	}
	fmt.Print("\n")

	var slice_2 []int = array[:3]
	for i, s := range slice_2 {
		fmt.Print(i, ", ")	// 0, 1, 2
		fmt.Print(s, "; ")	// -1, 2, 4
	}
	fmt.Print("\n")

	slice_2[0] = 100 	// slice change start array
	for i := 0; i < len(array); i++ {	
		fmt.Print(array[i], ", ")		// 100, 2, 4, 6, 8
	}
	fmt.Print("\n")

	// dynamic array
	// make([]T, length, capacity)
	dyn_array_five := make([]int, 5)
	dyn_array_five = append(dyn_array_five, 8, 16) 	// insert at the end
	fmt.Println(dyn_array_five) 					//  [0 0 0 0 0 8 16]

	// dynamic array 0
	dyn_array := make([]int, 0)
	dyn_array = append(dyn_array, 8, 16) 	// insert at the end
	fmt.Println(dyn_array) 					//  [8 16]

	for _, s := range dyn_array_five {
		fmt.Print(s, "; ")		// 0; 0; 0; 0; 0; 8; 16
	}
	fmt.Print("\n")

	word := "hello"
	for _, w := range word {
  		fmt.Printf("%c", w) 	// hello
	}
	fmt.Print("\n")

	sl_1 := make([]int, 10) // слайс ёмкостью (capacity) 10 и длиной (length) 10, заполненный нулями - [0 0 0 0 0 0 0 0 0 0]. 
	// И при добавлении элементов в слайс, будут добавляться после последнего 0, т.е. начиная с индекса 10. Повторное выделение памяти
	fmt.Println(sl_1)	// [0 0 0 0 0 0 0 0 0 0]
	sl_1 = append(sl_1, 1)
	fmt.Println(sl_1)	// [0 0 0 0 0 0 0 0 0 0 1] 

	sl_2 := make([]int, 0, 10) // будет создан слайс ёмкостью 10, но не содержащий ни одного элемента (длиной 0) - []. 
	// Добавление элементов в этот слайс начнётся с 0-го индекса. Повторное выделение памяти если превысит 10
	fmt.Println(sl_2)	// []
	sl_2 = append(sl_2, 1)
	fmt.Println(sl_2)	// [1] 

	sl_3 := make([]int, 5, 10)
	// будет создан слайс ёмкостью 10, но лежать в нём будут только 5 нулей (длина = 5)
	fmt.Println(sl_3)	// [0 0 0 0 0] 

	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"} // базовый массив

	users1 := initialUsers[2:6] 	// с 3-го по 6-й
	users2 := initialUsers[:4] 		// с 1-го по 4-й
	users3 := initialUsers[3:] 		// с 4-го до конца

	fmt.Println(users1) 		// [Kate Sam Tom Paul]
	fmt.Println(users2) 	// [Bob Alice Kate Sam]
	fmt.Println(users3) 	// [Sam Tom Paul Mike Robert]

	fmt.Printf(
		"Срез, основанный на базовом массиве длиной %d и емкостью %d: %v\n", 
		len(users1), 	// 4
		cap(users1),	// 6
		users1,			// [Kate Sam Tom Paul]
	)
	pointer := fmt.Sprintf("%p", users1)	// %p - адресс
	users1 = append(users1, "10")
	fmt.Printf("Массив: %v\n", initialUsers)	// Массив: [Bob Alice Kate Sam Tom Paul 10 Robert]
	fmt.Printf("Срез длиной %d и емкостью %d: %v\n", len(users1), cap(users1), users1)
	// Срез длиной 5 и емкостью 6: [Kate Sam Tom Paul 10]	
	fmt.Println(pointer == fmt.Sprintf("%p", users1)) 	// true
	fmt.Println("-------------------------------")
	users1 = append(users1, "11", "12", "13")
	fmt.Printf("Массив: %v\n", initialUsers) 	// Массив: [Bob Alice Kate Sam Tom Paul 10 Robert]
	fmt.Printf("Срез длиной %d и емкостью %d: %v\n", len(users1), cap(users1), users1)
	// Срез длиной 8 и емкостью 12: [Kate Sam Tom Paul 10 11 12 13]
	// Если емкости не достаточно, то создается новый срез, основанный на массиве большего объема
	fmt.Println(pointer == fmt.Sprintf("%p", users1))	// false

	delete := []int{1, 2, 3, 4, 5, 6, 7}
	delete = append(delete[0:2], delete[3:]...) // элементы передаются функции append как отдельные аргументы)
	fmt.Println(delete) 	// [1 2 4 5 6 7]

	// COPY
	fmt.Printf("delete = %v\n", delete)		// delete = [1 2 4 5 6 7]
	delete_new := make([]int, 3, 3)
	fmt.Printf("delete_new before copy = %v\n", delete_new)	// delete_new before copy = [0 0 0]
	count_element := copy(delete_new, delete) 
	fmt.Printf("delete_new after copy = %v\n", delete_new) 	// delete_new after copy = [1 2 4]
	fmt.Printf("Скопировано %d элемента\n", count_element) // Скопировано 3 элемента

	// function
	fnArray(array) 			// copy of array
	fnSlice(delete_new)		// link of array
	// fnArray(delete_new) 	// ERROR
	// fnSlice(array) 		// ERROR

	fmt.Println("array after function fnArray =", array) 			// [100 2 4 6 8]
	fmt.Println("delete_new after function fnSlice =", delete_new) 	// [1 15 4]
}