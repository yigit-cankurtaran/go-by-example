package main

import "fmt"

func zeroval(ival int)  {
  ival = 0
}

func zeroptr(iptr *int) {
  // *int = int pointer
  *iptr = 0
}

func main()  {
  i := 1
  fmt.Println("initial:", i)

  zeroval(i)
  fmt.Println("zeroval:", i)
  // doesn't change because it doesn't change what is stored in the memory

  zeroptr(&i)
  // &i = memory address of i
  fmt.Println("zeroptr:", i)

  fmt.Println("pointer:", &i)
}
