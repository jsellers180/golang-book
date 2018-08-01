package main

import "fmt"

func average(xs []float64) float64 {
    total := 0.0
    for _, v := range xs {
        total += v }
    return total / float64(len(xs))
  //panic("Not implemented")
}

func main() {
  /*
  og code
  var x [5]float64
  x[0] = 98
  x[1] = 93
  x[2] = 77
  x[3] = 82
  x[4] = 83
  var total float64 = 0
  for i := 0; i < len(x); i++ {
  total += x[i]
  }
  fmt.Println(total / float64(len(x)))*/
  // chap 7 code
  avg := []float64{98.3,93.1,77.0,82,73}
  fmt.Println(average(avg))
}
