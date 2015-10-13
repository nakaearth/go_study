package main

import (
  f "fmt"
) 

var (
a string = "test"
b        = "test2"
c        = "test3"
)

var i int

func hello() {
  f.Println("ほげ")
}

func plus_one(i int) int {
  return i + 1;
}

func main() {
  const Hello string = "Bye bye"

  //var message = "hello world"
  message := "hello world"

  f.Println(Hello)
  f.Println(message)
  f.Println(a)
  f.Println(b)
  f.Println(c)

  f.Println(i)

  for  n := 0 ; n < 10 ; n++ {
    f.Println(n)
    if n < 4 {
      f.Println("4よ小さい")
    } else if n >= 4  {
      f.Println("真ん中")
    } else {
      f.Println("その他")
    }
  }

  hello()

  f.Println(plus_one(4))
}
