package main
import "fmt"
func main () {
	// IF/ELSE
	var number_1, number_2 int
	fmt.Scan(&number_1)
	fmt.Scan(&number_2)
	if number_1 > number_2 { 		// открывающая скобка в строке с оператором
    	fmt.Println("number_1 > number_2")
	} else if number_1 == number_2 {
    	fmt.Println("number_1 = number_2")
	} else {
    	fmt.Println("number_1 < number_2")
	}

	if iq := 135; iq > 130 {
		fmt.Println("Genius!")
	} else {
		fmt.Println("not genius")
	}

	if iq := 100; iq > 130 {
		fmt.Println("Genius!")
	} else if (iq >= 90 && iq <= 130) {
		fmt.Println("normal")
	} else {
		fmt.Println("not genius")
	}
	// ERROR
	/*if iq := 135; iq > 130 {
		fmt.Println("Genius!")
	} 
	else {
		fmt.Println("not genius")
	}*/

	// SWITCH
	number := 2
	switch number {
    	case 1:
        	fmt.Println("One")
    	case 2:
        	fmt.Println("Two")
    	case 3:
        	fmt.Println("Three")
    	default:
        	fmt.Println(number)
	}

	month := "January"
	switch {		// после switch "переключатель" не нужен
    	case month == "December" || month == "January" || month == "February":
        	fmt.Println("Winter")
    	case month == "March" || month == "April" || month == "May":
        	fmt.Println("Spring")
		default:
        	fmt.Println(month)
	}
	
	day := 3
	switch day {
    	case 1:
        	fmt.Println("Monday")
        	fallthrough
    	case 2:
        	fmt.Println("Tuesday")
        	fallthrough
    	case 3:
        	fmt.Println("Wednesday")
        	fallthrough
    	case 4:
        	fmt.Println("Thursday")
        	fallthrough
    	case 5:
        	fmt.Println("Friday")
    	default:
        	fmt.Println("Error work day")
	}
	// FOR

	for count := 0; count <= 10; count++ {
		fmt.Println(count)
	}
	// ERROR
	//for count := 0; count <= 10; count++ 
	//{
	//}
	var i = 0
	for ; i < 10;{
    	fmt.Println(i)
    	i++
	}	
	// Analog while
	var i_1 = 0
	for i_1 < 10 {
		fmt.Println(i_1)
		i_1++
	}

	lang := ""
    fmt.Scanln(&lang)
    fmt.Println(lang)
}