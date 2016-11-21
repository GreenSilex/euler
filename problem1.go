/***
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.
***/

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
