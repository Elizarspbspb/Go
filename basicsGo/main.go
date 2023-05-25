package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    var x1 int = 1
    var x2, y1 = 2, 3
    x3, y2 := 4, 5                  // короткое объявление переменных
    fmt.Println("Var x1 = ", x1)    // Var x1 = 1
    fmt.Println("\tVar y1 = ", y1, "\nVar x2 = ", x2) //    Var y1 =  3 \n Var x2 =  2
    fmt.Println("Var y2 = ", y2, "Var x3 = ", x3)     // Var y2 =  5 Var x3 =  4

    var s = "Привет, Степик!"
    fmt.Println(s)          // Привет, Степик!
    fmt.Println(len(s))     // 27
    fmt.Println(utf8.RuneCountInString(s)) // 15
    fmt.Println(s[0], s[1]) // 208 159 

    var x4 int
    fmt.Println("Var x4 = ", x4) // Var x4 =  0
}
