/***
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
***/

package main

import "fmt"
import "time"
import "math"
import "runtime"

const(
    large_number = float64(600851475143)
    int_large_number = int64(large_number)
)

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

    runtime.GOMAXPROCS(2)   // set the number of CPUs to use
//    var low_num int64
//    var high_num int64

    time_first := time.Now()
    //high_num = int64(large_number/2)    // the highest number we need to check is half the large number
    //low_num = high_num - int64(math.Sqrt(large_number))   // the low numer to loop through

    // find the square root of the large number and turn it into an int
    max_erato_num := int64(math.Sqrt(large_number))
    erato_prime_slice := prime_numbers(max_erato_num)

    // let's get the largest prime factor
    answer := largest_prime_factor(erato_prime_slice, max_erato_num)

    //fmt.Println("erato slice: ", erato_prime_slice)   // if you want to see the prime numbers

    fmt.Println("length of slice", len(erato_prime_slice))

    fmt.Println("final answer: ", answer)
    time_elapsed := time.Since(time_first)
    fmt.Println(time_elapsed.Seconds())
}

func largest_prime_factor( erato_prime_slice []int64, max_erato_num int64) int64 {
    var i int64
    var j int64
    var answer int64
    var final_answer bool

    // we are testing to see if the large number is divisable by one of our primes, 
    // if it is, we are checking to see if that answer is a prime 
    for i = 0; i < int64(len(erato_prime_slice)); i++ {
        if int_large_number%erato_prime_slice[i] == 0 {
            answer = int_large_number/erato_prime_slice[i]
            fmt.Println("potential answer: ", answer)
            final_answer = true
            for j = 0; j <= int64(len(erato_prime_slice)); j++ {
                if answer%erato_prime_slice[j] == 0 {
                    fmt.Println("not: ", answer, "divisable by ", erato_prime_slice[j])
                    final_answer = false    // this number is not a prime
                    break
                }
            }
            // if our answer didn't get set to false, this should be the largest prime
            if final_answer {
                return answer
            }
        }
    }

    // if we got here it looks like one of our initial erato numbers is the largest prime factor
    for i = int64(len(erato_prime_slice)); i > 0; {
        i = i - 1
        if int_large_number%erato_prime_slice[i] == 0 {
            answer = erato_prime_slice[i]
            return answer
        }
    }
    return 0
}
