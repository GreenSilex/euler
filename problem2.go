package main

import "fmt"

func main() {
    // initial Fibonacci numbers and variable declaration
    all_fibonacci := []int{1,2}
    var fibonacci_length int
    even_fibonacci_total := 2
    var new_fibonacci_number int
    for {
        fibonacci_length = len(all_fibonacci)
        new_fibonacci_number = all_fibonacci[fibonacci_length-1] + all_fibonacci[fibonacci_length-2]
        // if the newest number is over 4 million, break out of loop
        if new_fibonacci_number > 4000000 {
            break
        }
        all_fibonacci = append(all_fibonacci, new_fibonacci_number)
        fmt.Println("fibonacci numbers: ", new_fibonacci_number)
        if new_fibonacci_number%2 == 0 {
        even_fibonacci_total = even_fibonacci_total + new_fibonacci_number
        }
        //fmt.Println("even_fibonacci number total: ", even_fibonacci_total)
    }
    fmt.Println("even_fibonacci number total: ", even_fibonacci_total)
}
