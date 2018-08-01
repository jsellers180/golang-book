package main

import "fmt"
// cheated
import "sort"

func main()  {
  x := []int{
      48,96,86,68,
      57,82,63,70,
      37,34,83,27,
      19,97, 9,17,
  }

  sort.Ints(x)
  lowest := x[0]
  fmt.Println(lowest)


/*
  for _,v:=range x {
    if v>n {
      fmt.Println(v,">",n)
    } else {
      fmt.Println(v,"<",n)
      n = v
      smallest = n
    }
  }
  fmt.Println("The smallest number is ", smallest)
*/
}
