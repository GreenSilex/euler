/***
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
***/

package main

import "fmt"

const(
  nth_prime_number = 10001
)

// This finds all the prime numbers up to the number you pass to the function
func prime_numbers(max_erato_num int64) ([]int64) {
    var i int64
    var j int64
    working_prime_slice := make([]int64, 1 )
    working_prime_map := make(map[int64]bool)
    working_prime_slice[0] = 2

    // find the primes in erato
    // build out the array with all the numbers
    for i = 3; i <= max_erato_num; i = i + 2 {
        working_prime_map[i] = true   // true == prime
    }
    fmt.Println("max erato num: ", max_erato_num)

    //fmt.Println("working_prime_map: ", working_prime_map)   //testing
    // loop through our erato map and mark non-primes as false
    for i = 3; i <= max_erato_num; i = i + 1 {
        if working_prime_map[i] == true {
            // loop through all the numbers changing non-primes to false
            j = i
            for j <= max_erato_num {
                j = j + i
                working_prime_map[j] = false
            }
        }
    }
    // create the erato prime slice
    for i = 3; i <= max_erato_num; i = i + 2 {
        if working_prime_map[i] == true {
            working_prime_slice = append(working_prime_slice, i)
        }
    }

    return working_prime_slice
}

func main() {
    rand_large_number := 120000   // choose some random large number to test primes up to
    prime_slice := prime_numbers(int64(rand_large_number))
    //fmt.Println(prime_slice)    // Testing print all primes
    fmt.Println("Number of primes in slice = ", len(prime_slice))
    fmt.Println("Answer for prime at position: ", nth_prime_number, " is: ", prime_slice[nth_prime_number-1])
}
