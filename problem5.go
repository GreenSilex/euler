/*** 
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
***/

package main

import "fmt"

func main() {

    is_answer := false

    // the number has to be divisiable by 20, so we should only count by 20's
    for i := 20; is_answer == false; i = i + 20 {
        if num_test(i) {
            fmt.Println("answer is: ", i)
            break
        }
    }
}


func num_test(potential_num int) bool {
    for i := 3; i < 20; i++ {
        if potential_num%i != 0 {
            return false
        }
    }
    // if we made it here, the number is divisiable by all digits from 1 - 20
    return true
}
