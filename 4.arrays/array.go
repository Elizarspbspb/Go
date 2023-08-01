package main

import "fmt"

func main () {
	words := [5]string{"Hello", "my", "name", "is", "Ravshan"} 
	for index, s := range words {
		fmt.Print(s, index, " ")  // Hello0 my1 name2 is3 Ravshan4 
	}
	fmt.Print("\n")
	for idx, _ := range words {
		fmt.Print(words[idx])		// HellomynameisRavshan 
	}
	fmt.Print("\n")
	for _, elem := range words {
		fmt.Print(elem)				// HellomynameisRavshan
	}
	fmt.Print("\n")
	fmt.Println(words) 				// [Hello my name is Ravshan]
	fmt.Print("\n")
	for i := 0; i < len(words); i++ {	
		fmt.Print(words[i], " ")		// Hello my name is Ravshan
	}
	fmt.Print("\n")
	var bl [5]bool
	fmt.Print(bl[0], bl[4])		// false false | true = 1, false = 0; 

	var a [3]int
	fmt.Println(a) 		// [0 0 0]

	var a1 [3]int = [3]int{1, 2, 3}
	b1 := [3]int{1, 2, 3}
	c1 := [...]int{1, 2, 3, 4}
	var in = 3
	//d1 := [4]int{1: 12, in: 3} 	// ERROR
	d1 := [4]int{1: 12, 3: in}

	fmt.Println(a1) 		// [1 2 3]
	fmt.Println(b1) 		// [1 2 3]
	fmt.Println(c1) 		// [1 2 3]
	fmt.Println(d1) 		// [0 12 0 3]

	fmt.Println(a1 == b1) 	// true
	fmt.Println(c1 == d1) 	// false

	for i := 0; i < len(d1); i++ {
		fmt.Printf("%d ", d1[i])
	}
	fmt.Println()
	for _, elem := range d1 {	// elem - только копия !
		//elem = 100
		fmt.Print(elem)		// 0 12 0 3
	}
	fmt.Println(d1) 			// [0 12 0 3]
	for idx := range d1 {
		d1[idx] = 100
		fmt.Println(d1[idx])	// 100
	}
	fmt.Println(d1) 			// [100 100 100 100]
	// Task 1: Ввести 10 чисел в массив, затем 3 пары индексов массива для изменения чисел на позициях 
	fmt.Println("-------- Task 1 --------")
	var ch1, ch2, temp uint8
	var workArray [10]uint8
	for i := 0; i < 10; i++ {
    	fmt.Scan(&workArray[i])
	}
	for i:=0;i<3;i++ {
    	fmt.Scan(&ch1, &ch2)
    	temp = workArray[ch1]
    	workArray[ch1] = workArray[ch2]
    	workArray[ch2] = temp
	}
	for i := 0; i < 10; i++ {
    	fmt.Printf("%d ", workArray[i])
	}
}