package main

import "fmt"

func main() {
  // slice1 defines a slice with 3 values
slice1 := []int{1,2,3}
// slice2 makes a slice with int data objects and room for 2 values
slice2 := make([]int, 2)
// this command copies the values of slice1 into slice2
copy(slice2, slice1)
// this will print slice1 as [1, 2, 3]
// this will print slice2 as [1, 2] because there is only room for 2 ints in the slice
fmt.Println(slice1, slice2)
}
