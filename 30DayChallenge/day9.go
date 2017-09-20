package main
import "fmt"

func factorial(x int) int{
    if x == 0 {
        return 1
    }
    return x * factorial(x-1)
}

func main() {
   //Enter your code here. Read input from STDIN. Print output to STDOUT
    var num int
    fmt.Scan(&num)

    fmt.Print(factorial(num))
}
