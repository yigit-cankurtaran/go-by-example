package main
import "fmt"

func main()  {
  var a [5]int
  fmt.Println("emp:", a)
  // prints 5 0's inside the array, which means by default an array is 0 valued.

  a[4] = 100
  // we don't use := because the variable is already declared
  fmt.Println("set:", a)
  fmt.Println("get:", a[4])

  fmt.Println("len:", len(a))

  b := [5]int{1, 2, 3, 4, 5}
  fmt.Println("dcl:", b)

  var twoD [2][3]int
  for i := 0; i < 2; i++ {
    for j := 0; j < 3; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2d: ", twoD)
}
