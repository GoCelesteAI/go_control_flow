package main

import "fmt"

func main() {
  hasLicense := true
  hasCar := true

  // AND operator (&&)
  if hasLicense && hasCar {
    fmt.Println("You can drive!")
  }

  // OR operator (||)
  isWeekend := false
  isHoliday := true
  if isWeekend || isHoliday {
    fmt.Println("No work today!")
  }

  // NOT operator (!)
  isRaining := false
  if !isRaining {
    fmt.Println("Great weather for a walk")
  }

  // Comparison operators
  x, y := 10, 20
  fmt.Println("x == y:", x == y)
  fmt.Println("x != y:", x != y)
  fmt.Println("x < y:", x < y)
}
