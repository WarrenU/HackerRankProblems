package main
import (
    "fmt"
    "math"
)

func hourglass_sum(arr [6][6]int) int {
    var max_sum = math.MinInt64
    var possible_max = 0

    for x := 0; x < 4; x++ {
        for y := 0; y < 4; y++ {
            possible_max = arr[x][y] + arr[x][y+1] + arr[x][y+2] + arr[x+1][y+1] + arr[x+2][y] + arr[x+2][y+1] + arr[x+2][y+2]
            if possible_max > max_sum {
                max_sum = possible_max
            }
        }
    }
    return max_sum
}

func main() {
    //Enter your code here. Read input from STDIN. Print output to STDOUT
    var data_arr = [6][6]int{}

    //Load in 2d array from std in:
    for i := 0; i < 6; i++ {
        for j := 0; j < 6; j++ {
            fmt.Scan(&data_arr[i][j])
        }
    }

    fmt.Print(hourglass_sum(data_arr))
}
