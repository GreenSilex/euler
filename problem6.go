/***
The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
***/

package main

import "fmt"
import "runtime"
import "sync"

func main() {
	runtime.GOMAXPROCS(2) // set the number of CPUs to use
	sum_square_chan := make(chan int64)
	square_sum_chan := make(chan int64)
	var wg sync.WaitGroup
	var sum_square int64
	var square_sum int64

	wg.Add(2)

	go sum_of_squares(sum_square_chan, &wg)
	go square_of_sums(square_sum_chan, &wg)
	sum_square = <-sum_square_chan
	square_sum = <-square_sum_chan

	wg.Wait()

	answer := square_sum - sum_square

	fmt.Println("sum of the squares: ", sum_square)
	fmt.Println("square of the sums: ", square_sum)
	fmt.Println("answer: ", answer)

}

func sum_of_squares(sum_square chan int64, wg *sync.WaitGroup) {
	defer wg.Done()

	var total int64
	var i int64

	for i = 1; i <= 100; i++ {
		total = total + i*i
	}

	sum_square <- total

}

func square_of_sums(square_sum chan int64, wg *sync.WaitGroup) {
	defer wg.Done()

	var total int64
	var i int64

	for i = 1; i <= 100; i++ {
		total = total + i
	}
	total = total * total
	square_sum <- total

}
