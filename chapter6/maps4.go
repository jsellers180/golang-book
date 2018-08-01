package main

import "fmt"

func main() {
  // creates the map name elements and sets it to map a string to a map. Then maps That
  // map to map a string to a string key pair
elements := map[string]map[string]string{
           "H": map[string]string{
                "name":"Hydrogen",
                "state":"gas",
                "number": "1",
           },
           "He": map[string]string{
                "name":"Helium",
                "state":"gas",
           },
           "Li": map[string]string{
                "name":"Lithium",
                "state":"solid",
           },
           "Be": map[string]string{
                "name":"Beryllium",
                "state":"solid",
           },
           "B":  map[string]string{
                "name":"Boron",
                "state":"solid",
           },
           "C":  map[string]string{
                "name":"Carbon",
                "state":"solid",
           },
           "N":  map[string]string{
                "name":"Nitrogen",
                "state":"gas",
           },
           "O":  map[string]string{
                "name":"Oxygen",
                "state":"gas",
           },
           "F":  map[string]string{
                "name":"Fluorine",
                "state":"gas",
           },
           "Ne":  map[string]string{
                "name":"Neon",
                "state":"gas",
           },
         }
     if el, ok := elements["H"]; ok {
           fmt.Println(el["name"], el["state"], el["number"])
         }
}
