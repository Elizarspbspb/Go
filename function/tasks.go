package main

import "fmt"

func max(num_1, num_2 int) (int) {
	if (num_1 >= num_2) {
		return num_1
	} else {
		return num_2
	}
}

func calc(num_1 int) (int, int) {
	return num_1*2, num_1*num_1
}

func isEven(num_1 int) (bool) {
	if ((num_1 % 2) == 0) {
		return true
	} else {
		return false
	}
}

func double(a int) {
    fmt.Println(a*2) 
} 

func double_m(num_1, num_2 int)(int) {
	count := 0
	if (num_1 < num_2) {
		for  ; num_1 <= num_2; num_1++ {
			count = count + num_1*num_1
		}
		return count
	} else {
		for  ; num_2 <= num_1; num_2++ {
			count = count + num_2*num_2
		}
		return count
	}
}

func main() {
	// Task 1: write function max() for 2 integer and return 1 max integer
	fmt.Println("-------- Task 1 --------")
	num_1, num_2 := 0, 0
	fmt.Scan(&num_1, &num_2)
	fmt.Println(max(num_1, num_2))

	// Task 2: write function calc() for 1 integer and return 2 - num*2 and num*num
	fmt.Println("-------- Task 2 --------")
	fmt.Println(calc(num_1))

	// Task 3: write function isEven() for 1 integer and return true or false if %2
	fmt.Println("-------- Task 3 --------")
	fmt.Println(isEven(num_1))

	// Task N
	for i := 4; i > 0; i-- {
        defer double(i)  // 2 4 6 8
    }

	// Task 4: write function double_m() for 2 integer and return sum number*number
	fmt.Println("-------- Task 4 --------")
	fmt.Println(double_m(num_1, num_2))
} 