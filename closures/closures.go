package main

import "fmt"

// closure = formed by anonymous functions.
func intSeq() func ()  int{
  i := 0
  return func() int{
    i++
    return i
  }
}

// the intSeq function wrapped the other nameless function inside itself

func main() {
  nextInt := intSeq()
  
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())

  newInts := intSeq()
  fmt.Println(newInts())
}
