package main
import "fmt"

func main() {
	var sumofsqrs int64 = 0
	var sqrofsum int64 = 0
	var i int64
	for i = 1; i <= 100; i++ {
		sumofsqrs += i*i
		sqrofsum += i
	}	
	sqrofsum *= sqrofsum
	fmt.Printf("The square of the sums is: %v \n", sqrofsum)
	fmt.Printf("The sum of the squares is: %v", sumofsqrs)
	fmt.Printf("The diference is: %v", sqrofsum - sumofsqrs)
}
