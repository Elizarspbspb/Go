package main

import "fmt"

func main() {

	var (
        number_1 int = 30
        number_2 int = 15
    )
    
	fmt.Println(number_1 == number_2) // false
	fmt.Println(number_1 != number_2) // true
	fmt.Println(number_1 > number_2)  // true
	fmt.Println(number_1 >= number_2) // true
	fmt.Println(number_1 < number_2)  // false
	fmt.Println(number_1 <= number_2) // false
    fmt.Println("------------")
	fmt.Println(number_1 != number_2 && number_1 >= number_2)  	// true
	fmt.Println(number_1 == number_2 || number_1 > number_2)   	// true
	fmt.Println(!(number_1 < number_2))          				// true
	fmt.Println("------------")	
	fmt.Println(!(8 >= 5) && 6 == 6) 		// output false
	fmt.Println(!(12 > 4 || 8 != 5)) 		// output false

	var a int = 8
    var b int = 3
    var c bool = a == b
    fmt.Println(a, b, c)      	// 8 3 false

	var a1 bool = true
	var b1 bool = !a1     		
	var c1 bool = !b1     		
	fmt.Println(a1, b1, c1)      		// true false true

	var b2 bool = 4 > 5 && 6 > 8   		
	var c2 bool = 3 <= 5 && 10 > 8 		
	fmt.Println(b2, c2)      			// false true

	var b3 bool = 4 > 5 || 6 > 8      	
	var c3 bool = 3 == 5 || 10 > 8    	
	fmt.Println(b3, c3)      			// false true

	fmt.Println("-------------------")
	fmt.Println(true||false)		// 0+1 = true
	fmt.Println(true&&false)		// 0*1 = false
	fmt.Println(true&&true)			// 1*1 = true
	fmt.Println(false||false)		// 0+0 = false
	fmt.Println(true||true)			// 1+1 = true
	fmt.Println(!true)				// ! 1 = false
}