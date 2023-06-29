package main

import "fmt"

type Contact struct {
    name string
    age  int
} 

func (x Contact) welcome() { // method
    fmt.Println(x.name)
    fmt.Println(x.age)
} 

func welcome(x Contact) {	// function with arguments = structur
    fmt.Println(x.name)
    fmt.Println(x.age)
}

func (ptr *Contact) increase(val int) {
    ptr.age += val
}

func main() {
	// obj := Contact // ERROR
	obj_1 := Contact{"Elizar", 26}
	var obj_2 = Contact{name: "Anton", age: 20}
	fmt.Println(obj_1.name, obj_2.name) // Elizar Anton

	point_e := &obj_1
	fmt.Println(point_e.age) 		// easy 26
	fmt.Println((*point_e).age) 	// full 26

	point_r := &Contact{"Roman", 23}
	point_r.welcome() 	// Roman 23
	obj_1.welcome() 	// Elizar 26
	
	welcome(obj_1) 		// Elizar 26

	obj_1.increase(4)
	fmt.Println(obj_1.age) // 30
}