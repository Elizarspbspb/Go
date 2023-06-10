package main

import "fmt"

var global = 0

func hello_stepik_1(a int, b int, c int) {
    fmt.Println("Hello, Stepik!", a + b + c)
}

func hello_stepik_2(a, b, c int){
    fmt.Println("Hello, Stepik!", a, b, c)
}
// ERROR
//func hello_stepik() 
//{
//    fmt.Println("Hello, Stepik!")
//}

func square(num float32) float32 { // int - type return variable
    return num * num
}

func concat(str1, str2 string) string { 
    str3 := str1+str2
    return str3
}

func plusMinus(num_1, num_2 int) (int, int) {
    return num_1 + num_2, num_1 - num_2
} 

func welcome() {
    fmt.Println("Welcome!")
}

func main() {
	num_1, num_2, num_3 := 1, 2, 3
    hello_stepik_1(num_1, num_2, num_3) // output: 1, 2, 3 or 6

	res := square(5.5)
    fmt.Println(res) // output: 30.25
    fmt.Println(square(6))  // output: 36

	line := concat("Hi", " baby")
	fmt.Println(line) // output: Hi baby

	addition, subtraction := plusMinus(15, 5)
	fmt.Println(addition, ":", subtraction) // 20 10
	fmt.Println(plusMinus(15, 5)) 			// 20 10

	defer welcome() 		// defer - only when finished main function !
    fmt.Println("Hello!")

	fmt.Println("Start")
    for i := 0; i < 5; i++ {
        defer fmt.Println(i)
    }
    fmt.Println("Finish") 	// Hello, Start, Finish, 4,3,2,1,0, Welcome
} 