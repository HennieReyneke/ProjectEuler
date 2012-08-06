// A concurrent prime sieve

package main

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int64) {
	var i int64
	for i = 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int64, out chan<- int64, prime int64) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func main() {
	var primesum int64
	ch := make(chan int64) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < 148933 ; i++ {
		var prime int64 
		prime = <-ch
		primesum += prime
		print("i: ", i, " prime: ", prime, " sum: ", primesum, "\n")
		ch1 := make(chan int64)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
		print("Prime sum: ", primesum)
}