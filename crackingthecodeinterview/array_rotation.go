/*
Given an array of n integers and a number, d,
 perform d left rotations on the array.
Then print the updated array as a single line of space-separated integers.
*/
package main
import (
    "fmt"
    "os"
    "bufio"
    "strings"
)


func rotate(x []string, num_rotations int) []string {
    return append(x[num_rotations:], x[:num_rotations]...)
}

func main() {
    //Enter your code here. Read input from STDIN. Print output to STDOUT
    var (
        arr_len int
        num_rotations int
    )
    fmt.Scan(&arr_len)
    fmt.Scan(&num_rotations)

    reader := bufio.NewReader(os.Stdin)
    arr_elements, _ := reader.ReadString('\n')

    var arr_to_rotate = strings.Split(arr_elements, " ")

    arr_to_rotate = rotate(arr_to_rotate, num_rotations)

    fmt.Print(strings.Join(arr_to_rotate, " "))
}
