package main
import "fmt"

func main() {
sumofeven := 2
	x := 1
	y := 2
	loop := true
for loop {
	x += y
	if x % 2 == 0 {
		if x <= 4000000 {
			sumofeven += x
		}else {
			loop = false
		 }
	}
	y += x
	if y % 2 == 0 {
		if y <= 4000000 {
			sumofeven += y
		}else {
			loop = false
		 }
	}
}
	fmt.Printf("The sumofeven is: %v", sumofeven)
}
