package main

import (
  "fmt"
  "log"
)

func main(){

  var userFloat float32
  
  // ask for user input
  fmt.Printf("Please enter a float: ")
  _, err :=
  fmt.Scan(&userFloat)
  
  // handle error
  if err != nil {
    log.Fatal(err)
  }

  // convert float to int
  var userInt int = int(userFloat)
  fmt.Printf("The truncate value of %f is %d", userFloat, userInt)
}
