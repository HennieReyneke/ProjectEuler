package main
import "fmt"

func main() {
	var prime int64 = 0
	var factorcnt int64 = 0
	var i int64
	var j int64
	for i = 30; i < 300425737571; i++ {
	for j = 2; j < i/2; j++ {
		if i % j == 0 {
			factorcnt++
		}	
	}
	if factorcnt == 0 && 600851475143 % i == 0 {
		prime = i
		fmt.Printf("Prime: %v \n", prime)		
	}
	factorcnt = 0
	}
	fmt.Printf("Highest Prime is: %v", prime)
}
