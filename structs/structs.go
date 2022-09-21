package main

import "fmt"

type person struct {
  name string
  age int
}

func newPerson(name string) *person {
  // you can safely return a pointer to a local variable

  // p := person {name: name}
  // if you do it that way name becomes whatever you pass it to
  p := person {name: "Jen"}
  // if you do it this way the name becomes whatever was passed in this line, it doesn't change
  p.age = 42
  return &p
}

func main() {
  fmt.Println(person{"Bob", 20})
  fmt.Println(person{name: "Alice", age: 30})
  // both of them working
  fmt.Println(person{name: "Fred"})
  // default is 0 or space
  fmt.Println(&person{name: "Alice", age: 30})
  // pointer thingy didnt work?
  fmt.Println(newPerson("Jon"))

  s := person{name: "Sean", age: 50}
  fmt.Println(s.name)

  sp := &s
  // this is the s pointer
  // struct pointers get automatically dereferenced
  fmt.Println(sp.age)

  sp.age = 51
  fmt.Println(sp.age)
  // it changes, so structs are mutable
}
