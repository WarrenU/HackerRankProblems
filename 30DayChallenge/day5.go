package main
import (
    "fmt"
)

func main() {
    //Enter your code here. Read input from STDIN. Print output to STDOUT
    var multiplier int
    fmt.Scan(&multiplier)
    for i := 1; i <= 10; i++ {
        fmt.Printf("%d x %d = %d\n", multiplier, i, multiplier*i)
    }
}
