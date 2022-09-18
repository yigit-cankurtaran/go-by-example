package main

import "fmt"

func fact(n int) int  {
  if n == 0 {
    return 1
  }
  return n * fact (n - 1)
}

func main()  {
  fmt.Println(fact(7))
  // 1, 2, 6, 24, 120, 720, 5040

  var fib func (n int) int

  fib = func(n int) int {
    if n < 2 {
      return n
    }

    return fib(n - 1) + fib(n - 2)
  }
  // normal fibonacci olayı

  fmt.Println(fib(7))
}
