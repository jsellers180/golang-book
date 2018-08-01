package main

import "fmt"

// no floats, the decimal value is just cut off
func main() {
  var x [5]int
  x[0] = 98
  x[1] = 93
  x[2] = 77
  x[3] = 82
  x[4] = 83
  var total = 0
  for _, value := range x {
    total += value
}
  /*for i := 0; i < len(x); i++ {
  total += x[i]
  }*/
  fmt.Println(total / (len(x)))
}
