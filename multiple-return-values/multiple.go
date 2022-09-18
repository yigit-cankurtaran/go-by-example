package main

import "fmt"

func vals() (int, int) {
  return 3, 7
}

func main()  {
  a, b := vals()
  fmt.Println(a)
  fmt.Println(b)

  _, c := vals()
  fmt.Println(c)
  // you can't declare something without using it, so if you only want one thing you gotta declare it using the blank identifier "_"

  d, _ := vals()
  fmt.Println(d)
}
