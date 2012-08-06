package main
import (
	"fmt"
	)
func main() {
//n := 21
var a [21][21]int64
a[0][0] = 0
for i := 1; i < 21; i++ {
a[0][i] = 1
a[i][0] = 1
}
for x := 1; x < 21; x++{
for y := 1; y < 21; y++{
a[x][y] = a[x-1][y] + a[x][y-1] 
}
}
fmt.Printf("The number of routes are: %v", a[20][20])
}

