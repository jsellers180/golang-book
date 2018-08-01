package main

import "fmt"

func main() {
  bool := (true && false) || (false && true) || !(false && false)
 fmt.Println( bool )
 fmt.Println(true && false)
 fmt.Println(true || true)
 fmt.Println(true || false)
 fmt.Println(!true)
}
