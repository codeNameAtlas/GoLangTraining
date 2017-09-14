package main

import (

  "fmt"
)

const yardToMeters float64 = 1.09361

func main() {

  fmt.Print("Enter your workout in meters :")
  var meters float64
  //scan is pointing the above text to a place in memory
  fmt.Scan(&meters)

  yards := meters * yardToMeters
  
  fmt.Println(meters, "meters is ", yards, "many yeards")


}
