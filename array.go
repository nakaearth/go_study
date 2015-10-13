
package main

import (
  f "fmt"
)

var arr1 = [4]string{"a",  "b",  "c",  "d"}
var arr2 = [...]string{"a",  "b",  "c",  "d"}


func main() {
  f.Println(arr1)
  f.Println(arr2)

  var s = []string{"a",  "b",  "c",  "d"}

  f.Println(s[0])
}
