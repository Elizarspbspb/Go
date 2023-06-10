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

func one_or_two(num_1 int, num_2 int, str string) (int, string) {
	if (str == "one") {
		return num_1, str
	} else if (str == "two") {
		return num_2, str
	} else {
		return 0, str
	}
}

func mars_age(num_1 int) (int) {
	num_1 = num_1 * 365 / 687
	return num_1
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

	// Task 5: write function one_or_two() for 2 integer and 1 string (one or two) and return string and 1 num if one , 2 num if two , 0 if not one and not two
	fmt.Println("-------- Task 5 --------")
	fmt.Println(one_or_two(num_1, num_2, "one"))

	// Task 6: write function mars_age() for 1 integer Age on Earth and return 1 integer Age on Mars
	fmt.Println("-------- Task 6 --------")
	fmt.Println(mars_age(num_1))
} 