package main

import "fmt"

func main() {
	fmt.Println("hello world")
	type Celsius float64
	var temp Celsius = 10
	fmt.Println(temp)

	var x int = 34
	var y int
	var ip *int // ip is a pointer
  
  ip = &x
  y = *ip

  fmt.Println(ip)
  fmt.Println(y)
  
  var name string = "Dima"
  fmt.Printf("My name is %s", name)
  fmt.Println("")
  fmt.Printf("Temperature is %f", temp)
	
}


