package main

import "fmt"

func main() {
  // Short statement: declare in if
  if num := 10; num%2 == 0 {
    fmt.Println(num, "is even")
  } else {
    fmt.Println(num, "is odd")
  }

  // Variable scoped to if block
  if length := len("Hello"); length > 3 {
    fmt.Println("String has", length, "chars")
  }

  // Combine with conditions
  if x := 15; x > 10 && x < 20 {
    fmt.Println(x, "is between 10 and 20")
  }
}
