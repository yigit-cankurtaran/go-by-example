package main

import "fmt"

func main()  {
  nums := []int{2, 3, 4}
  sum := 0

  for _, num := range nums{
    sum += num
  }
  // if we don't use the blank identifier "_" it just gives us how many numbers there are

  fmt.Println("sum:", sum)

  for i, num := range nums{
    if num == 3{
      fmt.Println("index:", i)
    }
  }

  kvs := map[string]string{"a": "apple", "b": "banana"}
  for k, v := range kvs{
    fmt.Printf("%s -> %s \n", k, v)
  }

  for k := range kvs {
    fmt.Println("key:", k)
  }

  // this part of the code uses the unicode code points
  for i, c := range "go" {
    fmt.Println(i, c)
  }
}
