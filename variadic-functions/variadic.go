package main

import "fmt"

func sum(nums ...int)  {
  fmt.Print(nums, " ")
  // print something, put the 2nd thing between those
  total := 0

  for _, num := range nums {
    // if you don't add the _ it adds an extra step?? weird
    total += num
  }
  fmt.Println(total)
}

func main()  {
  sum (1, 2)
  sum (1, 2, 3)

  nums := []int{1, 2, 3, 4}
  sum (nums...)

  nums = []int{1, 2, 3, 4, 5}
  sum (nums...)
}
