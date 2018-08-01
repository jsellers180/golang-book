package main

import "fmt"

// takes the total of two arrays and finds the win percent
func main() {
	fmt.Print("How many wins? ")
	var input string
	fmt.Scanln(&input)
	fmt.Println("You had " + input + " wins")
  win := []float64{9,1,4,7,9,3,2,6,6}
  loss := []float64{4,3,6,5,4,2,3,5,3}
  winTotal := 0.0
  lossTotal := 0.0
    for _, v := range win {
        winTotal += v }

    for _, v := range loss {
        lossTotal += v }
    matchTotal := winTotal + lossTotal
  fmt.Println(winTotal / matchTotal)
}
