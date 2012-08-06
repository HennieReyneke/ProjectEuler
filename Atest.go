package main
import (
    "fmt"
    "strconv"
    "strings"
    )
func main() {
    s := "876567747896354336739443262"
    var dstr []string = strings.SplitAfterN(s,"",len(s))
    var dint [27]int
    for i := 0; i < len(s); i++ {
        dint[i], _ = strconv.Atoi(dstr[i]) //index out of range at runtime
        fmt.Printf("dstr[%v] is: %s\n", i, dstr[i])
        fmt.Printf("dint[%v] is: %v\n", i, dint[i])
    }
}
