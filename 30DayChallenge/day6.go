package main
import "fmt"

func main() {
    //Enter your code here. Read input from STDIN. Print output to STDOUT
    var (
        num_of_inputs int
        str_to_parse string
    )
    fmt.Scan(&num_of_inputs)
    for i := 0; i < num_of_inputs; i++ {
        fmt.Scan(&str_to_parse)
        var (
            even_idx_str string
            odd_idx_str string
        )
        for i := 0; i < len(str_to_parse); i++ {
            if i%2 == 0 {
                even_idx_str += string(str_to_parse[i])
            } else {
                odd_idx_str += string(str_to_parse[i])
            }
        }
        fmt.Printf("%s %s\n", even_idx_str, odd_idx_str)
    }
}
