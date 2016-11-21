/***
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
***/
package main

import "fmt"
import "strconv"
import "strings"

func main() {

    answer := num_loop()

    fmt.Println("answer: ", answer)
}

func num_loop() int {

    numA := 999
    numB := 999
    var potential_answer int
    answer := 0

    for numA > 99 {
        numB = 999
        for numB >= numA {
            potential_answer = numB * numA
            // fmt.Println(numB, " * ", numA, " = ", potential_answer)    // testing
            if palindrome(potential_answer) {
                // our looping method doesn't garantee we are testing the highest product numbers first
                // so test the numbers and return only the highest palindrome
                if potential_answer > answer {
                    answer = potential_answer
                }
            }
            numB = numB - 1
        }
        numA = numA - 1
    }
    return answer
}


func palindrome (potential_answer int) bool {

    start := 0

    potential_answer_string := strconv.Itoa(potential_answer)
    potential_answer_length := len(potential_answer_string)
    potential_answer_split := strings.Split(potential_answer_string, "")
    end := potential_answer_length - 1

    for start < potential_answer_length/2 {

        if potential_answer_split[start] == potential_answer_split[end] {
            start = start + 1
            end = end - 1
        } else {
            // not a palindrome, return false for the next number
            return false
        }
    }
    // we tested all the digits, we must have a palindrome
    return true
}

