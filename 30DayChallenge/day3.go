package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    //Enter your code here. Read input from STDIN. Print output to STDOUT
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    num, _ := strconv.ParseInt(text, 10, 64)

    if num%2 == 1 {
        fmt.Print("Weird")
    }else{
        if 2 <= num && num <= 5 {
            fmt.Print("Not Weird")
        }else if 6 <= num && num <= 20 {
            fmt.Print("Weird")
        }else if num > 20 {
            fmt.Print("Not Weird")
        }
    }
}
