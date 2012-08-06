package main
import "fmt"

func main() {
	loop := true
	for x := 1; x < 1000 && loop; x++ {
	for y := 1; y < 1000 && loop; y++ {
	for z := 1; z < 1000 && loop; z++ {
		if x*x + y*y == z*z &&
			x + y + z == 1000 {
			fmt.Printf("X is: %v", x)
			fmt.Printf("Y is: %v", y)
			fmt.Printf("Z is: %v", z)
			fmt.Printf("Answer is: %v", x*y*z)
			loop = false
		}
	}
	}
	}
}
