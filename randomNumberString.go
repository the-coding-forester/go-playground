package main

import (
  "fmt"
  "math/rand"
)
/*
By convention, the package name is the same as the last element of the import path.
For instance, the "math/rand" package comprises files that begin with the statement package rand.
*/

func main(){
  fmt.Println("My favorite number is", rand.Intn(10))
  fmt.Println("My favorite number is", rand.Intn(10))
  fmt.Println("My favorite number is", rand.Intn(10))
  fmt.Println("My favorite number is", rand.Intn(10))
}
