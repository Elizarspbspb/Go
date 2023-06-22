package main

import "fmt"

func main () {
	// map, dictionaries, associative arrays or hash tables
	map_1 := make(map[string]int) 	// map, key - string, value - int
	map_1["James"] = 42
	map_1["Amy"] = 24
	fmt.Println(map_1["James"]) 	//  42	

	map_2 := map[string]int{
		"James": 42, 
		"Amy": 24 }
	fmt.Println(map_2["James"]) 	//  42	

	delete(map_2, "James")
	map_2["Bob"] = 8
	fmt.Println(map_2)	// map[Amy:24 Bob:8]

	numbers := map[int]int{
		8: 2, 
		2: 5,
		4: 10,
		5: 4,
	}    
	fmt.Println(numbers)	// sort - map[2:5 4:10 5:4 8:2]
	delete(numbers, 2)   
	fmt.Println(numbers)	// map[4:10 5:4 8:2] 
	fmt.Println(numbers[4] - numbers[5]) 	// 6
}