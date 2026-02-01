package main

import "fmt"

func main() {
  age := 15

  // if-else
  if age >= 18 {
    fmt.Println("You can vote")
  } else {
    fmt.Println("You cannot vote yet")
  }

  // else-if chain
  score := 75
  if score >= 90 {
    fmt.Println("Grade: A")
  } else if score >= 80 {
    fmt.Println("Grade: B")
  } else if score >= 70 {
    fmt.Println("Grade: C")
  } else {
    fmt.Println("Grade: F")
  }
}
