/*
Given a base-10 integer, d, convert it to binary (base-2).
 Then find and print the base-10 integer denoting the maximum number of consecutive 1's in d's binary representation.
*/

package main
import (
    "fmt"
    "strconv"
)

func main() {
   //Enter your code here. Read input from STDIN. Print output to STDOUT
    var d int64
    fmt.Scan(&d)

    var binary_rep = strconv.FormatInt(d, 2)
    var max_consecutive_ones int = 0
    var curr_sequence_of_ones int = 0
    for _, e := range binary_rep {
        if (string(e) == "1") {
            curr_sequence_of_ones += 1
        }
        if (curr_sequence_of_ones > max_consecutive_ones) {
            max_consecutive_ones = curr_sequence_of_ones
        }
        if (string(e) == "0") {
            curr_sequence_of_ones = 0
        }
    }
    fmt.Print(max_consecutive_ones)
}
