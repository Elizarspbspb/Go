package main

import "fmt"

func div(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("делитель равен 0")
    }
    return a / b, nil
}

func foo() {
    // паникуем
    panic("unexpected!")
}

func main() {
    d, err := div(10, 0)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("d = %d", d)
    }

	// выполняется при завершении main
    defer func() {
        // вызываем recover и сравниваем результат с nil
        if r := recover(); r != nil {
            fmt.Println(r) // выведет "unexpected"
        }
    }()
    foo() // внутри foo срабатывает паника
    fmt.Println("Вы не увидите это сообщение, так как в foo возникла паника")
}