package main
import (
	"fmt"
	"strconv"
	"strings"
	)
func main() {
	s := "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"
	var dstr []string = strings.SplitAfterN(s,"",1000)
	//var dint []int
	i0 := 0
	i1 := 0
	i2 := 0
	i3 := 0
	i4 := 0
	result := 0
	n := []byte(s)
	//var d []int
	for x := 0; x < 1000; x++ {
		dint[x], _ := strconv.Atoi(dstr[x]) 
		fmt.Printf("dint[%v] is: %v\n", x, dstr[x])
		}	
	for i := 0; i < 996; i++ {
			n0 := n[i]
			n1 := n[i+1]
			n2 := n[i+2]
			n3 := n[i+3]
			n4 := n[i+4]
			ni0 := int(n0)
			ni1 := int(n1)
			ni2 := int(n2)
			ni3 := int(n3)
			ni4 := int(n4)
			if ni0*ni1*ni2*ni3*ni4 > result {
				result = ni0*ni1*ni2*ni3*ni4
				i0 = ni0
				i1 = ni1
				i2 = ni2
				i3 = ni3
				i4 = ni4
	fmt.Printf("i0 is: %v\n", i0)
	fmt.Printf("i1 is: %v\n", i1)
	fmt.Printf("i2 is: %v\n", i2)
	fmt.Printf("i3 is: %v\n", i3)
	fmt.Printf("i4 is: %v\n", i4)			}
	}
	fmt.Printf("i0 is: %v\n", i0)
	fmt.Printf("i1 is: %v\n", i1)
	fmt.Printf("i2 is: %v\n", i2)
	fmt.Printf("i3 is: %v\n", i3)
	fmt.Printf("i4 is: %v\n", i4)
	fmt.Printf("Answer is: %v", result)
}
