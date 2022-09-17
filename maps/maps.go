package main
import "fmt"

func main()  {
  // hash or dict in other languages
  m := make(map[string]int)

  m["k1"] = 7
  m["k2"] = 13

  fmt.Println("map:", m)

  v1 := m["k1"]
  fmt.Println("v1:", v1)

  fmt.Println("len:", len(m))

  delete(m, "k2")
  fmt.Println("After delete:")
  fmt.Println("map:", m)

  _, prs := m["k2"]
  // prs is used to check if the key we want exists, _ is the blank identifier because we didn't need the value itself
  fmt.Println("prs:", prs)

  n := map[string]int{"foo": 1, "bar": 2}
  // you can use this syntax to initialize a map in 1 line instead of declaring different keys in different lines like we did above
  fmt.Println("map:", n)
  fmt.Printf("%v", n) // you use %v to print maps with printf
}
