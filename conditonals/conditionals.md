# Conditionals in Go

In Go, conditionals are used to execute different blocks of code based on certain conditions. The primary conditional statements in Go are `if`, `else if`, and `else`. Here is a brief overview:

## `if` Statement
The `if` statement evaluates a condition and executes the block of code inside it if the condition is true.

```go
if condition {
    // code to be executed if condition is true
}
```

## `else if` Statement
The `else if` statement allows you to check multiple conditions. If the first condition is false, it checks the next condition.

```go
if condition1 {
    // code to be executed if condition1 is true
} else if condition2 {
    // code to be executed if condition2 is true
}
```

## `else` Statement
The `else` statement is executed if none of the preceding conditions are true.

```go
if condition1 {
    // code to be executed if condition1 is true
} else if condition2 {
    // code to be executed if condition2 is true
} else {
    // code to be executed if none of the above conditions are true
}
```

## Example
Here is an example of using conditionals in Go:

```go
package main

import "fmt"

func main() {
    x := 10

    if x < 0 {
        fmt.Println("x is negative")
    } else if x == 0 {
        fmt.Println("x is zero")
    } else {
        fmt.Println("x is positive")
    }
}
```

In this example, the program checks if `x` is negative, zero, or positive and prints the corresponding message.

## Short Variable Declaration in `if`
Go allows you to declare and initialize a variable in the `if` statement itself, which is scoped to the `if` block.

```go
if y := 42; y > 0 {
    fmt.Println("y is positive")
}
```

This feature is useful for concise and readable code, especially when the variable is only needed within the conditional block.