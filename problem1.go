package main

import "fmt"

func main() {
    top_number := 1000
    total_added := 0

    // loop through 1 to 999
    for i := 0; i < top_number; i++ {
        three_remainder := i%3
        five_remainder := i%5
        if three_remainder == 0 {
            total_added = total_added+i
        } else if five_remainder == 0 {
             total_added = total_added+i
        }
        fmt.Println(total_added)
    }
}
