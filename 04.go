package main
import (
	"fmt"
	"strconv"
	)
func main() {
	palindrome := 0
	loop := true
	for i := 998; loop; i-- {
		s := strconv.Itoa(i)
		s += s
		a := []byte(s)
		a[5] = a[0]
		a[4] = a[1]
		a[3] = a[2]
		palistr := string(a)
		palindrome, _ = strconv.Atoi(palistr)
	for j := 998; j > 400; j-- {
		if palindrome % j == 0 && palindrome / j < 1000 {
			fmt.Printf("The palindrome is: %v \n", palindrome)
			fmt.Printf("Factor1 is: %v \n", j)
			fmt.Printf("Factor2 is: %v \n", palindrome / j)
			loop = false
			break
		}	
	}
	}
}
