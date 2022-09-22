package main

import "fmt"

func main()  {
  for i := 1; i < 100; i++{
    var a bool = i % 3 == 0
    var b bool = i % 5 == 0

    if a && b {
      fmt.Println("Fizzbuzz")
    } else if a {
      fmt.Println("Fizz")
    } else if b {
      fmt.Println("Buzz")
    } else {
      fmt.Println(i)
    }
  }
}
