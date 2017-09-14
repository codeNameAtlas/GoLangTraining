//an example where i pass a pointer to another function and change the value

package main

import "fmt"

//func zero is expecting a reference to the pointer as *int
func zero(x *int) {
  fmt.Println("x before we change the value",x)
  *x = 0
}

func main() {

  x := 5
  fmt.Println("value of x :", x)
  fmt.Println("memory of x :", &x)
  //pass memory of x into a different function
  zero(&x)
  fmt.Println("value of x now inside main", x)
  fmt.Println("memory of x now inside main", &x)
}
