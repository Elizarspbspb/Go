package main

import "fmt"

func main () {
	words := [5]string{"Hello", "my", "name", "is", "Ravshan"} 
	for i, s := range words {
		fmt.Print(s, i, " ")  // Hello0 my1 name2 is3 Ravshan4 
	}
	fmt.Print("\n")
	fmt.Println(words) // [Hello my name is Ravshan]
	fmt.Print("\n")
	for i := 0; i < len(words); i++ {	
		fmt.Print(words[i], " ")		// Hello my name is Ravshan
	}
}