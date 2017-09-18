package main
import (
    "os"
    "fmt"
    "bufio"
    "strings"
)

func main() {
    //Enter your code here. Read input from STDIN. Print output to STDOUT
    var array_length int
    fmt.Scan(&array_length)

    scanner := bufio.NewScanner(os.Stdin)

    var inpt string
    for scanner.Scan() {
        inpt = scanner.Text()
    }
    var ordered_vals []string = strings.Split(inpt," ")

    for i := array_length-1; i > -1; i-- {
        fmt.Printf("%s ", ordered_vals[i])
    }
}
