package main

import (
    "fmt"
)

// iota - идентификатор Go для упрощения определений увеличивающихся чисел
const(
	c0 = iota  // c0 == 0
	c1 = iota  // c1 == 1
	c2 = iota  // c2 == 2
)

const(
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	_  // пропускаем 7
	Add
)

const(
    A2 int = 45
    B2 // скопирует А
    C2 float32 = 3.3
    D2 // скопирует С
    Sunday2 = iota 	// iota = 4
    Monday2
    Tuesday2
    Wednesday2
    Thursday2
    Friday2
    Saturday2
)

const (
	a = iota + 1	// iota = 0
	_				// iota = 1
	b				// iota = 2
	c				// iota = 3
	d = c + 2		// iota = 4
	t				// iota = 5
	i				// iota = 6
	i2 = iota + 2	// iota = 7; i2 = 7 + 2 = 9
)

func main() {
	fmt.Println(c0, c1, c2) // вывод: 0 1 2

	fmt.Println(Sunday)   	// вывод: 0
	fmt.Println(Saturday) 	// вывод: 6
	fmt.Println(Add) 		// вывод: 8

	// переменные нет ни в одном из описанных блоков const, поэтому индекс не увеличился
	const x = iota  	// x == 0
	const y = iota  	// y == 0
	fmt.Println(x, y) 	//  0 0

	fmt.Println(Friday2)	// 9

	fmt.Println(i2)		// 9
}
