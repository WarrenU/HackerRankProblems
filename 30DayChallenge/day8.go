/*
Task
Given  names and phone numbers, assemble a phone book that maps
friends' names to their respective phone numbers.
You will then be given an unknown number of names to query your phone book for.
 For each  queried, print the associated entry from your phone book
  on a new line in the form name=phoneNumber;
   if an entry for  is not found, print Not found instead.
*/

package main
import (
    "os"
    "bufio"
    "fmt"
)

func main() {
   //Enter your code here. Read input from STDIN. Print output to STDOUT
    var (
        map_len int
        input_key string
        input_value string
    )

    m := make(map[string]string)

    fmt.Scan(&map_len)

    for i := 0; i < map_len; i++ {
        fmt.Scan(&input_key)
        fmt.Scan(&input_value)
        m[input_key] = input_value
    }

    var query_input string
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        query_input = scanner.Text()
        if val, ok := m[query_input]; ok{
            fmt.Printf("%s=%s\n", query_input, val)
        } else {
            fmt.Println("Not found")
        }
    }
}
