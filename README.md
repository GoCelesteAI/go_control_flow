# Go Control Flow: if/else

Code examples from **Go Tutorial for Beginners #6 - Control Flow: if/else**

## Contents

- `basic_if.go` - Basic if statements
- `if_else.go` - if-else and else-if chains
- `short_statement.go` - Short statement syntax (declare variables in if)
- `logical_operators.go` - Logical operators (&&, ||, !) and comparison operators

## Running the Examples

```bash
go run basic_if.go
go run if_else.go
go run short_statement.go
go run logical_operators.go
```

## Key Concepts

### Basic if Statement
```go
if age >= 18 {
    fmt.Println("You are an adult")
}
```

### if-else Chain
```go
if score >= 90 {
    fmt.Println("Grade: A")
} else if score >= 80 {
    fmt.Println("Grade: B")
} else {
    fmt.Println("Grade: F")
}
```

### Short Statement Syntax
```go
if num := 10; num%2 == 0 {
    fmt.Println(num, "is even")
}
```

### Logical Operators
```go
if hasLicense && hasCar {
    fmt.Println("You can drive!")
}
```

## Watch the Tutorial

[Go Tutorial for Beginners #6 - Control Flow: if/else](https://youtube.com/@CelesteAI)

## License

MIT
