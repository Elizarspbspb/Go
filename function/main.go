package main

import "fmt"

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

func main() {
	num_1, num_2, num_3 := 1, 2, 3
    hello_stepik_1(num_1, num_2, num_3) // output: 1, 2, 3 or 6
} 