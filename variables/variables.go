package main
import "fmt"

func main()  {
  var a = "initial"
  fmt.Println(a)

  var b, c int = 1, 2
  // in go you declare it as "var age int"
  fmt.Println(b, c)

  var d = true
  fmt.Println(d)

  var e int
  fmt.Println(e)
  // int default is 0

  f := "apple"
  // ":=" is shorthand for declaring and initializing a variable
  // f := apple == f string = apple
  fmt.Println(f)
}
