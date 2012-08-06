package main
import "fmt"

func main() {
	divcnt := 0
	for i := 20; true; i++ {
	for j := 2; j <= 20; j++ {
		if i % j == 0 {
			divcnt++
		}	
	}
	if divcnt == 19 {
		fmt.Printf("Answer is: %v", i)
		break
	}
	divcnt = 0
	}
}
