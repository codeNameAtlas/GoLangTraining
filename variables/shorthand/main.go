package main

import "fmt"

func main() {
//shorthand must be declared in an executable funcion
  a := 10
  b := "string"
  c := true
  d := 4.17

  fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

//to print the types
  fmt.Printf("%T \n", a)
  fmt.Printf("%T \n", b)
  fmt.Printf("%T \n", c)
  fmt.Printf("%T \n", d)

//zero value

  var e int
  var f float64
  var g string
  var h bool

  fmt.Printf("%v \n", e)
  fmt.Printf("%v \n", f)
  fmt.Printf("%v \n", g)
  fmt.Printf("%v \n", h)



}
