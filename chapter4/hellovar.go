package main

import "fmt"

func main() {
  fmt.Print("Enter the temp in F: ")
  var input float64
  fmt.Scanf("%f", &input)
  output := (input - 32) * 5/9
  fmt.Println("The temp in C is", output)
}

/* func second() {
  fmt.Println(x, "dibs")
}
*/
