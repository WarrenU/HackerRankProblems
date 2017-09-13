package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
  "math"
)

func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}

func main(){
    scanner := bufio.NewScanner(os.Stdin)

    var meal_cost float64
    var tip_percentage uint64
    var tax_percentage uint64
    var total_cost float64

    value := make([]string, 3)
    var count int = 0
    for scanner.Scan() {
      value[count] = scanner.Text()
      count++
    }

    meal_cost, _ = strconv.ParseFloat(value[0], 64)
    tip_percentage, _ = strconv.ParseUint(value[1], 10, 64)
    tax_percentage, _ = strconv.ParseUint(value[2], 10, 64)

    var tip float64 = meal_cost * (float64(tip_percentage)/100.0)
    //fmt.Printf("%v\n", tip)
    var tax float64 = meal_cost * (float64(tax_percentage)/100.0)
    //fmt.Printf("%v\n", tax)
    total_cost = meal_cost + tip + tax

    fmt.Printf("The total meal cost is %d dollars.", round(total_cost))
}
