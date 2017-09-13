package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  var i uint64 = 4
  var d float64 = 4.0
  var s string = "HackerRank "

  scanner := bufio.NewScanner(os.Stdin)
  // Declare second integer, double, and String variables.
  var i2 uint64
  var d2 float64
  var s2 string

  // Read and save an integer, double, and String to your variables.
  value := make([]string, 3)
  var count int = 0
  for scanner.Scan() {
    value[count] = scanner.Text()
    count++
  }
  i2, _ = strconv.ParseUint(value[0], 10, 64)
  d2, _ = strconv.ParseFloat(value[1], 64)
  s2 = value[2]

  // Print the sum of both integer variables on a new line.
  fmt.Printf("%d\n", i+i2)

  // Print the sum of the double variables on a new line.
  fmt.Printf("%.1f\n", d+d2)

  // Concatenate and print the String variables on a new line
  // The 's' variable above should be printed first.
  fmt.Printf("%v", s+s2)
}
