package main

import (

  "fmt"
)

const p string = "constant named string"

func main() {

  const s = "constant no name"

  fmt.Println("p :", p)
  fmt.Println("s :", s)
}

//types are checking at time of compiling. statically compiled
//a "kind", the compiler will determine at the time of compiling what the value
//of an unnamed constant is
