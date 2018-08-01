package main

import "fmt"

func main() {
  // maps a string to an int
  x := make(map[string]int)
  // sets x["key"] to 10
  x["key"] = 10
  // x["key"] is now accessible
  fmt.Println(x["key"])
}
