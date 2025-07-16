# Go Phase 1 Mastery Guide: Fundamentals Deep Dive

## 1.1 Language Basics

### Variables and Types

#### Basic Types Deep Dive

Go has a rich type system with specific characteristics you must understand:

**Numeric Types:**
```go
// Signed integers
var i8 int8 = 127      // -128 to 127
var i16 int16 = 32767  // -32,768 to 32,767
var i32 int32 = 2147483647
var i64 int64 = 9223372036854775807

// Unsigned integers
var ui8 uint8 = 255    // 0 to 255 (also byte)
var ui16 uint16 = 65535
var ui32 uint32 = 4294967295
var ui64 uint64 = 18446744073709551615

// Machine-dependent types
var i int = 42        // int32 or int64 depending on platform
var ui uint = 42      // uint32 or uint64 depending on platform
var ptr uintptr = 42  // holds pointer addresses

// Floating-point
var f32 float32 = 3.14159
var f64 float64 = 3.141592653589793

// Complex numbers
var c64 complex64 = 1 + 2i
var c128 complex128 = 1 + 2i
```

**String and Boolean:**
```go
// Strings are immutable UTF-8 byte sequences
var str string = "Hello, 世界"
var empty string    // "" (empty string)

// Booleans
var flag bool = true
var defaultBool bool // false
```

**Critical Understanding Points:**
- Types are NOT interchangeable: `int32` ≠ `int64` ≠ `int`
- No automatic type conversion (unlike C/Java)
- `rune` is an alias for `int32` (represents Unicode code points)
- `byte` is an alias for `uint8`

#### Type Inference with `:=`

The short variable declaration is powerful but has rules:

```go
func typeInferenceExamples() {
    // Basic inference
    name := "John"        // string
    age := 25            // int
    height := 5.9        // float64
    isActive := true     // bool
    
    // Multiple assignment
    x, y := 10, 20
    a, b, c := 1, 2.5, "hello"
    
    // With function returns
    result, err := someFunction()
    
    // Redeclaration rules
    name = "Jane"        // OK: assignment
    name := "Bob"        // ERROR: redeclaration
    
    // Mixed declaration and assignment
    name, surname := "John", "Doe"  // name reassigned, surname declared
}
```

**Key Rules:**
- `:=` can only be used inside functions
- At least one variable must be new
- Type is inferred from right-hand side
- Default to `int`, `float64`, `string`, `bool`

#### Constants and iota

Constants are compile-time values with unique features:

```go
// Basic constants
const Pi = 3.14159
const Name = "Go"
const Active = true

// Typed constants
const TypedInt int = 42
const TypedFloat float64 = 3.14

// Untyped constants (more flexible)
const UntypedInt = 42      // can be used as any integer type
const UntypedFloat = 3.14  // can be used as any float type

// Constant expressions (evaluated at compile time)
const (
    SecondsPerMinute = 60
    MinutesPerHour   = 60
    SecondsPerHour   = SecondsPerMinute * MinutesPerHour
)
```

**iota - The Constant Generator:**
```go
// Basic iota usage
const (
    Sunday = iota    // 0
    Monday          // 1
    Tuesday         // 2
    Wednesday       // 3
    Thursday        // 4
    Friday          // 5
    Saturday        // 6
)

// Advanced iota patterns
const (
    _ = iota                    // skip 0
    KB = 1 << (10 * iota)      // 1024
    MB                         // 1048576
    GB                         // 1073741824
    TB                         // 1099511627776
)

// Multiple constants per line
const (
    A, B = iota, iota << 10    // 0, 0
    C, D                       // 1, 1024
    E, F                       // 2, 2048
)

// Expressions with iota
const (
    FileRead = 1 << iota       // 1
    FileWrite                  // 2
    FileExecute                // 4
    FileReadWrite = FileRead | FileWrite  // 3
)
```

#### Zero Values - The Foundation

Understanding zero values is crucial for Go mastery:

```go
func zeroValues() {
    var i int         // 0
    var f float64     // 0.0
    var b bool        // false
    var s string      // ""
    var ptr *int      // nil
    var slice []int   // nil
    var m map[string]int // nil
    var ch chan int   // nil
    var fn func()     // nil
    var iface interface{} // nil
    
    // Structs get zero values for all fields
    type Person struct {
        Name string
        Age  int
    }
    var p Person      // {Name: "", Age: 0}
    
    // Arrays get zero values for all elements
    var arr [5]int    // [0, 0, 0, 0, 0]
}
```

**Zero Value Best Practices:**
- Design your types to be useful with zero values
- Use zero values for initialization when possible
- Understand nil vs empty for reference types

### Control Structures

#### if/else Statements

Go's if statements have unique features:

```go
func ifStatements() {
    x := 10
    
    // Basic if
    if x > 5 {
        fmt.Println("x is greater than 5")
    }
    
    // If with else
    if x > 15 {
        fmt.Println("x is greater than 15")
    } else {
        fmt.Println("x is not greater than 15")
    }
    
    // If with else if
    if x > 20 {
        fmt.Println("large")
    } else if x > 10 {
        fmt.Println("medium")
    } else {
        fmt.Println("small")
    }
    
    // If with initialization (powerful feature)
    if y := x * 2; y > 15 {
        fmt.Println("y is", y)
        // y is only accessible within this if block
    }
    
    // Common pattern with error handling
    if err := someFunction(); err != nil {
        log.Fatal(err)
    }
    
    // Multiple initializations
    if a, b := getValue(); a > b {
        fmt.Println("a is greater")
    }
}
```

#### for Loops - The Only Loop

Go has only one looping construct but it's versatile:

```go
func forLoops() {
    // Traditional for loop
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }
    
    // While-style loop
    i := 0
    for i < 10 {
        fmt.Println(i)
        i++
    }
    
    // Infinite loop
    for {
        // Do something forever
        break // Exit condition
    }
    
    // Range over slice
    numbers := []int{1, 2, 3, 4, 5}
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
    
    // Range over map
    person := map[string]int{"Alice": 25, "Bob": 30}
    for name, age := range person {
        fmt.Printf("%s is %d years old\n", name, age)
    }
    
    // Range over string (runes)
    for index, runeValue := range "Hello, 世界" {
        fmt.Printf("Index: %d, Rune: %c\n", index, runeValue)
    }
    
    // Ignore index with blank identifier
    for _, value := range numbers {
        fmt.Println(value)
    }
    
    // Only index
    for index := range numbers {
        fmt.Println(index)
    }
}
```

**Advanced for Loop Patterns:**
```go
func advancedForPatterns() {
    // Nested loops with labels
    outer:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i == 1 && j == 1 {
                break outer  // Break out of outer loop
            }
            fmt.Printf("i=%d, j=%d\n", i, j)
        }
    }
    
    // Continue with labels
    outer2:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 1 {
                continue outer2  // Continue outer loop
            }
            fmt.Printf("i=%d, j=%d\n", i, j)
        }
    }
}
```

#### switch Statements

Go's switch is more powerful than in other languages:

```go
func switchStatements() {
    day := 3
    
    // Basic switch
    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    default:
        fmt.Println("Other day")
    }
    
    // Multiple values per case
    switch day {
    case 1, 2, 3, 4, 5:
        fmt.Println("Weekday")
    case 6, 7:
        fmt.Println("Weekend")
    }
    
    // Switch with expressions
    switch {
    case day < 1 || day > 7:
        fmt.Println("Invalid day")
    case day <= 5:
        fmt.Println("Weekday")
    default:
        fmt.Println("Weekend")
    }
    
    // Switch with initialization
    switch today := time.Now().Weekday(); today {
    case time.Saturday, time.Sunday:
        fmt.Println("Weekend")
    default:
        fmt.Println("Weekday")
    }
    
    // Type switch (interface{})
    var x interface{} = 42
    switch v := x.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

**Switch Best Practices:**
- No `break` needed (automatic)
- Use `fallthrough` for C-style behavior (rarely needed)
- Expression switches are more idiomatic than if-else chains
- Type switches are powerful for interface handling

#### defer Statements

`defer` is a unique Go feature for cleanup:

```go
func deferStatements() {
    // Basic defer
    defer fmt.Println("This executes last")
    fmt.Println("This executes first")
    
    // File handling with defer
    file, err := os.Open("example.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()  // Guaranteed to execute
    
    // Multiple defers (LIFO order)
    defer fmt.Println("First defer")
    defer fmt.Println("Second defer")
    defer fmt.Println("Third defer")
    // Output: Third defer, Second defer, First defer
    
    // Defer with arguments evaluation
    x := 10
    defer fmt.Println("x was", x)  // x evaluated immediately (10)
    x = 20
    // Will print "x was 10", not "x was 20"
    
    // Defer with function literals
    defer func() {
        fmt.Println("x is now", x)  // x evaluated at defer time (20)
    }()
}

// Defer for resource management
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()  // Always closes, even if panic occurs
    
    // Read file operations...
    return nil
}

// Defer for mutex unlocking
func criticalSection() {
    mu.Lock()
    defer mu.Unlock()  // Guaranteed unlock
    
    // Critical section code...
}
```

### Functions

#### Function Declaration and Calling

Go functions have a clean syntax with powerful features:

```go
// Basic function
func add(a int, b int) int {
    return a + b
}

// Same type parameters (shorthand)
func multiply(a, b int) int {
    return a * b
}

// Multiple parameters of different types
func describe(name string, age int, active bool) {
    fmt.Printf("%s is %d years old and active: %t\n", name, age, active)
}

// No parameters, no return
func sayHello() {
    fmt.Println("Hello!")
}

// Calling functions
func functionCalls() {
    result := add(5, 3)
    fmt.Println(result)
    
    product := multiply(4, 7)
    fmt.Println(product)
    
    describe("Alice", 25, true)
    sayHello()
}
```

#### Multiple Return Values

Go's multiple return values are idiomatic:

```go
// Function returning multiple values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Function returning multiple values of same type
func getNameAndSurname() (string, string) {
    return "John", "Doe"
}

// Using multiple returns
func multipleReturns() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Result: %.2f\n", result)
    }
    
    // Ignoring returns with blank identifier
    result, _ = divide(10, 2)  // Ignore error
    _, err = divide(10, 0)     // Ignore result
    
    // Multiple assignment
    name, surname := getNameAndSurname()
    fmt.Printf("Full name: %s %s\n", name, surname)
}

// Common pattern: value and boolean
func getValue(key string) (string, bool) {
    m := map[string]string{"name": "John", "age": "25"}
    value, exists := m[key]
    return value, exists
}
```

#### Named Return Values

Named returns can make code more readable:

```go
// Named return values
func calculate(a, b int) (sum, product int) {
    sum = a + b
    product = a * b
    return  // Naked return
}

// More complex example
func processData(data []int) (processed []int, count int, err error) {
    if len(data) == 0 {
        err = errors.New("empty data")
        return
    }
    
    processed = make([]int, 0, len(data))
    for _, v := range data {
        if v > 0 {
            processed = append(processed, v*2)
            count++
        }
    }
    return
}

// Named returns with explicit return
func divide2(a, b float64) (result float64, err error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    result = a / b
    return result, nil  // Explicit return
}
```

**Named Return Best Practices:**
- Use for complex functions with multiple returns
- Be careful with naked returns in long functions
- Named returns are zero-initialized

#### Variadic Functions

Functions that accept variable number of arguments:

```go
// Basic variadic function
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Variadic with other parameters
func greet(greeting string, names ...string) {
    for _, name := range names {
        fmt.Printf("%s, %s!\n", greeting, name)
    }
}

// Using variadic functions
func variadicExamples() {
    fmt.Println(sum(1, 2, 3, 4, 5))  // 15
    fmt.Println(sum(10, 20))         // 30
    fmt.Println(sum())               // 0
    
    greet("Hello", "Alice", "Bob", "Charlie")
    
    // Passing slice to variadic function
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println(sum(numbers...))  // Spread operator
}

// fmt.Printf is a variadic function
func printfExamples() {
    fmt.Printf("Name: %s, Age: %d\n", "John", 25)
    fmt.Printf("Just a string\n")
}
```

#### Anonymous Functions and Closures

Go supports first-class functions:

```go
func anonymousAndClosures() {
    // Anonymous function
    func() {
        fmt.Println("Anonymous function executed")
    }()
    
    // Anonymous function with parameters
    result := func(a, b int) int {
        return a + b
    }(5, 3)
    fmt.Println(result)  // 8
    
    // Assigning anonymous function to variable
    multiply := func(a, b int) int {
        return a * b
    }
    fmt.Println(multiply(4, 5))  // 20
    
    // Closure capturing variables
    x := 10
    increment := func() {
        x++  // Captures x from outer scope
    }
    
    fmt.Println(x)  // 10
    increment()
    fmt.Println(x)  // 11
    
    // Closure factory
    makeCounter := func() func() int {
        count := 0
        return func() int {
            count++
            return count
        }
    }
    
    counter1 := makeCounter()
    counter2 := makeCounter()
    
    fmt.Println(counter1())  // 1
    fmt.Println(counter1())  // 2
    fmt.Println(counter2())  // 1
}

// Higher-order functions
func forEach(slice []int, fn func(int)) {
    for _, v := range slice {
        fn(v)
    }
}

func filter(slice []int, predicate func(int) bool) []int {
    var result []int
    for _, v := range slice {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}

func higherOrderExamples() {
    numbers := []int{1, 2, 3, 4, 5}
    
    // Using forEach
    forEach(numbers, func(n int) {
        fmt.Printf("%d ", n)
    })
    
    // Using filter
    evens := filter(numbers, func(n int) bool {
        return n%2 == 0
    })
    fmt.Println(evens)  // [2, 4]
}
```

## 1.2 Data Structures

### Arrays and Slices

#### Arrays - Fixed Size Collections

Arrays in Go are values, not references:

```go
func arrayBasics() {
    // Declaration and initialization
    var arr1 [5]int  // [0, 0, 0, 0, 0]
    
    // Initialize with values
    arr2 := [5]int{1, 2, 3, 4, 5}
    
    // Let compiler count elements
    arr3 := [...]int{1, 2, 3, 4, 5}
    
    // Sparse initialization
    arr4 := [5]int{1: 10, 3: 30}  // [0, 10, 0, 30, 0]
    
    // Access and modify
    arr1[0] = 42
    fmt.Println(arr1[0])  // 42
    
    // Length is part of type
    fmt.Println(len(arr2))  // 5
    
    // Arrays are values (copied)
    arr5 := arr2  // Copies all elements
    arr5[0] = 999
    fmt.Println(arr2[0])  // Still 1
    fmt.Println(arr5[0])  // 999
}

// Arrays as function parameters (copied)
func modifyArray(arr [5]int) {
    arr[0] = 100  // Doesn't affect original
}

// Arrays by reference
func modifyArrayByRef(arr *[5]int) {
    arr[0] = 100  // Modifies original
}
```

#### Slices - Dynamic Arrays

Slices are the primary array-like structure in Go:

```go
func sliceBasics() {
    // Creating slices
    var slice1 []int  // nil slice
    slice2 := []int{1, 2, 3, 4, 5}
    slice3 := make([]int, 5)      // [0, 0, 0, 0, 0]
    slice4 := make([]int, 5, 10)  // len=5, cap=10
    
    // Slicing arrays or slices
    arr := [5]int{1, 2, 3, 4, 5}
    slice5 := arr[1:4]  // [2, 3, 4]
    slice6 := arr[:3]   // [1, 2, 3]
    slice7 := arr[2:]   // [3, 4, 5]
    slice8 := arr[:]    // [1, 2, 3, 4, 5]
    
    fmt.Println(slice1 == nil)  // true
    fmt.Println(len(slice2))    // 5
    fmt.Println(cap(slice2))    // 5
}

// Append operations
func sliceAppend() {
    var slice []int
    fmt.Printf("len=%d, cap=%d\n", len(slice), cap(slice))  // 0, 0
    
    slice = append(slice, 1)
    fmt.Printf("len=%d, cap=%d\n", len(slice), cap(slice))  // 1, 1
    
    slice = append(slice, 2, 3, 4)
    fmt.Printf("len=%d, cap=%d\n", len(slice), cap(slice))  // 4, 4
    
    slice = append(slice, 5)
    fmt.Printf("len=%d, cap=%d\n", len(slice), cap(slice))  // 5, 8 (capacity doubled)
    
    // Append slice to slice
    slice2 := []int{6, 7, 8}
    slice = append(slice, slice2...)
    fmt.Println(slice)  // [1, 2, 3, 4, 5, 6, 7, 8]
}
```

#### Slice Internals and Memory Management

Understanding slice internals is crucial:

```go
func sliceInternals() {
    // Slice header contains: pointer, length, capacity
    arr := [5]int{1, 2, 3, 4, 5}
    slice := arr[1:4]  // Points to arr[1], len=3, cap=4
    
    // Modifying slice affects underlying array
    slice[0] = 100
    fmt.Println(arr)    // [1, 100, 3, 4, 5]
    fmt.Println(slice)  // [100, 3, 4]
    
    // Appending within capacity
    slice = append(slice, 999)
    fmt.Println(arr)    // [1, 100, 3, 4, 999]
    fmt.Println(slice)  // [100, 3, 4, 999]
    
    // Appending beyond capacity creates new array
    slice = append(slice, 888)
    fmt.Println(arr)    // [1, 100, 3, 4, 999] (unchanged)
    fmt.Println(slice)  // [100, 3, 4, 999, 888] (new array)
}

// Copy function
func sliceCopy() {
    source := []int{1, 2, 3, 4, 5}
    
    // Create destination slice
    dest := make([]int, len(source))
    
    // Copy elements
    copied := copy(dest, source)
    fmt.Printf("Copied %d elements\n", copied)  // 5
    
    // Modify destination
    dest[0] = 999
    fmt.Println(source)  // [1, 2, 3, 4, 5] (unchanged)
    fmt.Println(dest)    // [999, 2, 3, 4, 5]
    
    // Partial copy
    dest2 := make([]int, 3)
    copy(dest2, source)
    fmt.Println(dest2)  // [1, 2, 3]
}
```

#### Multi-dimensional Slices

Creating and working with 2D slices:

```go
func multiDimensionalSlices() {
    // 2D slice
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    
    // Access elements
    fmt.Println(matrix[1][2])  // 6
    
    // Create dynamic 2D slice
    rows, cols := 3, 4
    matrix2 := make([][]int, rows)
    for i := range matrix2 {
        matrix2[i] = make([]int, cols)
    }
    
    // Fill with values
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            matrix2[i][j] = i*cols + j
        }
    }
    
    // Print matrix
    for _, row := range matrix2 {
        fmt.Println(row)
    }
}
```

### Maps

Maps are Go's built-in associative data type:

```go
func mapBasics() {
    // Declaration and initialization
    var map1 map[string]int  // nil map
    map2 := make(map[string]int)
    map3 := map[string]int{
        "apple":  5,
        "banana": 3,
        "orange": 8,
    }
    
    // Adding elements
    map2["key1"] = 10
    map2["key2"] = 20
    
    // Reading elements
    value := map3["apple"]
    fmt.Println(value)  // 5
    
    // Safe reading with existence check
    value, exists := map3["grape"]
    if exists {
        fmt.Printf("grape: %d\n", value)
    } else {
        fmt.Println("grape not found")
    }
    
    // Deleting elements
    delete(map3, "banana")
    fmt.Println(map3)  // map[apple:5 orange:8]
    
    // Map length
    fmt.Println(len(map3))  // 2
}

// Map iteration
func mapIteration() {
    scores := map[string]int{
        "Alice": 95,
        "Bob":   87,
        "Carol": 92,
    }
    
    // Iterate over key-value pairs
    for name, score := range scores {
        fmt.Printf("%s: %d\n", name, score)
    }
    
    // Iterate over keys only
    for name := range scores {
        fmt.Printf("Name: %s\n", name)
    }
    
    // Iterate over values only
    for _, score := range scores {
        fmt.Printf("Score: %d\n", score)
    }
}

// Maps as reference types
func mapReference() {
    original := map[string]int{"a": 1, "b": 2}
    
    // Assignment creates reference, not copy
    copy := original
    copy["c"] = 3
    
    fmt.Println(original)  // map[a:1 b:2 c:3]
    fmt.Println(copy)      // map[a:1 b:2 c:3]
    
    // To create actual copy
    actualCopy := make(map[string]int)
    for k, v := range original {
        actualCopy[k] = v
    }
    
    actualCopy["d"] = 4
    fmt.Println(original)   // map[a:1 b:2 c:3]
    fmt.Println(actualCopy) // map[a:1 b:2 c:3 d:4]
}
```

### Structs

Structs are Go's way to create custom types:

```go
// Basic struct definition
type Person struct {
    Name string
    Age  int
    Email string
}

// Struct with embedded types
type Address struct {
    Street string
    City   string
    State  string
    Zip    string
}

type Employee struct {
    Person  // Embedded struct
    Address // Embedded struct
    ID      int
    Salary  float64
}

func structBasics() {
    // Struct creation
    var p1 Person  // Zero value struct
    fmt.Println(p1)  // {  0}
    
    // Struct literal
    p2 := Person{
        Name:  "Alice",
        Age:   30,
        Email: "alice@example.com",
    }
    
    // Positional initialization (not recommended)
    p3 := Person{"Bob", 25, "bob@example.com"}
    
    // Partial initialization
    p4 := Person{Name: "Charlie"}  // Age: 0, Email: ""
    
    // Accessing fields
    fmt.Println(p2.Name)  // Alice
    p2.Age = 31
    fmt.Println(p2.Age)   // 31
}

// Struct methods
func (p Person) String() string {
    return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

func (p Person) IsAdult() bool {
    return p.Age >= 18
}

// Pointer receiver (can modify struct)
func (p *Person) HaveBirthday() {
    p.Age++
}

// Value receiver (cannot modify struct)
func (p Person) GetInfo() string {
    return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func structMethods() {
    p := Person{Name: "David", Age: 17}
    
    fmt.Println(p.String())     // David (17 years old)
    fmt.Println(p.IsAdult())    // false
    
    p.HaveBirthday()
    fmt.Println(p.Age)          // 18
    fmt.Println(p.IsAdult())    // true
    
    info := p.GetInfo()
    fmt.Println(info)
}
```

#### Embedded Structs (Composition)

Go uses composition instead of inheritance:

```go
// Base types
type Address struct {
    Street string
    City   string
    State  string
    Zip    string
}

type Person struct {
    Name  string
    Age   int
    Email string
}

// Composition through embedding
type Employee struct {
    Person          // Embedded struct
    Address         // Embedded struct
    ID       int
    Salary   float64
    Department string
}

func embeddedStructs() {
    emp := Employee{
        Person: Person{
            Name:  "Alice",
            Age:   30,
            Email: "alice@company.com",
        },
        Address: Address{
            Street: "123 Main St",
            City:   "Anytown",
            State:  "CA",
            Zip:    "12345",
        },
        ID:         1001,
        Salary:     75000.00,
        Department: "Engineering",
    }
    
    // Access embedded fields directly
    fmt.Println(emp.Name)    // Alice (from Person)
    fmt.Println(emp.Street)  // 123 Main St (from Address)
    fmt.Println(emp.ID)      // 1001 (from Employee)
    
    // Access embedded structs explicitly
    fmt.Println(emp.Person.Name)   // Alice
    fmt.Println(emp.Address.City)  // Anytown
    
    // Methods are promoted from embedded types
    fmt.Println(emp.IsAdult())     // true (from Person)
    emp.HaveBirthday()             // calls Person's method
    fmt.Println(emp.Age)           // 31
}

// Method promotion with embedded structs
func (a Address) FullAddress() string {
    return fmt.Sprintf("%s, %s, %s %s", a.Street, a.City, a.State, a.Zip)
}

func methodPromotion() {
    emp := Employee{
        Person: Person{Name: "Bob", Age: 25},
        Address: Address{
            Street: "456 Oak Ave",
            City:   "Springfield",
            State:  "IL",
            Zip:    "62701",
        },
        ID: 1002,
    }
    
    // Method from embedded Address is promoted
    fmt.Println(emp.FullAddress())  // 456 Oak Ave, Springfield, IL 62701
}

// Handling method conflicts
type Manager struct {
    Employee
    TeamSize int
}

// Override embedded method
func (m Manager) GetInfo() string {
    return fmt.Sprintf("Manager: %s, Team Size: %d", m.Name, m.TeamSize)
}

func methodOverride() {
    mgr := Manager{
        Employee: Employee{
            Person: Person{Name: "Carol", Age: 35},
            ID:     2001,
        },
        TeamSize: 5,
    }
    
    // Uses Manager's GetInfo method
    fmt.Println(mgr.GetInfo())  // Manager: Carol, Team Size: 5
    
    // Access Employee's GetInfo explicitly
    fmt.Println(mgr.Employee.GetInfo())  // Name: Carol, Age: 35
}
```

#### Struct Tags for JSON/XML

Struct tags provide metadata for serialization:

```go
type User struct {
    ID       int    `json:"id" xml:"id"`
    Name     string `json:"name" xml:"name"`
    Email    string `json:"email" xml:"email"`
    Password string `json:"-" xml:"-"`                    // Ignored
    Age      int    `json:"age,omitempty" xml:"age"`      // Omit if zero
    Active   bool   `json:"is_active" xml:"is_active"`
    Created  time.Time `json:"created_at" xml:"created_at"`
}

func structTags() {
    user := User{
        ID:       1,
        Name:     "Alice",
        Email:    "alice@example.com",
        Password: "secret123",
        Age:      0,  // Will be omitted due to omitempty
        Active:   true,
        Created:  time.Now(),
    }
    
    // JSON serialization
    jsonData, err := json.Marshal(user)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(jsonData))
    
    // JSON deserialization
    jsonStr := `{"id":2,"name":"Bob","email":"bob@example.com","is_active":false}`
    var newUser User
    err = json.Unmarshal([]byte(jsonStr), &newUser)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%+v\n", newUser)
}

// Advanced struct tags
type Product struct {
    ID          int     `json:"id" db:"product_id" validate:"required"`
    Name        string  `json:"name" db:"product_name" validate:"required,min=3,max=50"`
    Price       float64 `json:"price" db:"price" validate:"required,gt=0"`
    Description string  `json:"description,omitempty" db:"description"`
    InStock     bool    `json:"in_stock" db:"in_stock"`
}

// Custom struct tag parsing
func parseStructTags() {
    t := reflect.TypeOf(Product{})
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        jsonTag := field.Tag.Get("json")
        dbTag := field.Tag.Get("db")
        validateTag := field.Tag.Get("validate")
        
        fmt.Printf("Field: %s\n", field.Name)
        fmt.Printf("  JSON: %s\n", jsonTag)
        fmt.Printf("  DB: %s\n", dbTag)
        fmt.Printf("  Validate: %s\n", validateTag)
        fmt.Println()
    }
}
```

## 1.3 Pointers and Memory

### Pointers

Pointers store memory addresses:

```go
func pointerBasics() {
    // Basic pointer operations
    x := 42
    var p *int = &x  // p points to x
    
    fmt.Println(x)   // 42
    fmt.Println(&x)  // Memory address of x
    fmt.Println(p)   // Memory address of x
    fmt.Println(*p)  // 42 (dereferencing)
    
    // Modify through pointer
    *p = 100
    fmt.Println(x)   // 100
    
    // Zero value of pointer is nil
    var ptr *int
    fmt.Println(ptr == nil)  // true
    
    // Cannot dereference nil pointer
    // fmt.Println(*ptr)  // panic: runtime error
}

// Pointers with structs
func pointerStructs() {
    type Person struct {
        Name string
        Age  int
    }
    
    p := Person{Name: "Alice", Age: 30}
    
    // Pointer to struct
    ptr := &p
    
    // Access fields through pointer
    fmt.Println(ptr.Name)  // Alice (automatic dereferencing)
    fmt.Println((*ptr).Name)  // Alice (explicit dereferencing)
    
    // Modify through pointer
    ptr.Age = 31
    fmt.Println(p.Age)  // 31
}

// Pointer receivers in methods
func pointerReceivers() {
    type Counter struct {
        count int
    }
    
    // Value receiver - cannot modify
    func (c Counter) GetCount() int {
        return c.count
    }
    
    // Pointer receiver - can modify
    func (c *Counter) Increment() {
        c.count++
    }
    
    // Pointer receiver - cannot modify if method doesn't
    func (c *Counter) Reset() {
        c.count = 0
    }
    
    counter := Counter{count: 5}
    fmt.Println(counter.GetCount())  // 5
    
    counter.Increment()  // Go automatically takes address
    fmt.Println(counter.GetCount())  // 6
    
    // Explicit pointer usage
    ptr := &counter
    ptr.Increment()
    fmt.Println(counter.GetCount())  // 7
}
```

#### Pointer Arithmetic (Limited in Go)

Go has limited pointer arithmetic for safety:

```go
func pointerArithmetic() {
    // Go does NOT support pointer arithmetic like C/C++
    // This is for memory safety
    
    arr := [5]int{1, 2, 3, 4, 5}
    ptr := &arr[0]
    
    // This would be invalid in Go:
    // ptr++  // Error: invalid operation
    // ptr += 1  // Error: invalid operation
    
    // Instead, use indices
    fmt.Println(arr[0])  // 1
    fmt.Println(arr[1])  // 2
    
    // For advanced pointer operations, use unsafe package
    // (generally not recommended)
    import "unsafe"
    
    ptr1 := &arr[0]
    ptr2 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr1)) + unsafe.Sizeof(int(0))))
    fmt.Println(*ptr2)  // 2 (second element)
}
```

#### Pointers vs Values in Function Parameters

Understanding when to use pointers vs values:

```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Value receiver - receives copy
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Pointer receiver - receives reference
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// Function with value parameter
func calculateArea(r Rectangle) float64 {
    r.Width = 999  // Doesn't affect original
    return r.Width * r.Height
}

// Function with pointer parameter
func scaleRectangle(r *Rectangle, factor float64) {
    r.Width *= factor   // Affects original
    r.Height *= factor
}

func pointersVsValues() {
    rect := Rectangle{Width: 10, Height: 5}
    
    // Value method
    fmt.Println(rect.Area())  // 50
    
    // Pointer method
    rect.Scale(2)
    fmt.Println(rect.Area())  // 200 (10*2 * 5*2)
    
    // Function with value parameter
    area := calculateArea(rect)
    fmt.Println(area)         // 9990 (999 * 10)
    fmt.Println(rect.Width)   // 20 (unchanged)
    
    // Function with pointer parameter
    scaleRectangle(&rect, 0.5)
    fmt.Println(rect.Width)   // 10 (changed)
}

// When to use pointers vs values
func pointerGuidelines() {
    // Use pointers when:
    // 1. You need to modify the receiver
    // 2. The receiver is large (expensive to copy)
    // 3. Consistency (if some methods use pointers, all should)
    
    // Use values when:
    // 1. The receiver is small (maps, slices, interfaces, channels)
    // 2. The receiver is immutable
    // 3. The receiver is a basic type or small struct
}
```

#### Memory Allocation with new and make

Go provides two built-in functions for memory allocation:

```go
func memoryAllocation() {
    // new() - allocates memory and returns pointer
    p1 := new(int)
    fmt.Println(p1)   // Memory address
    fmt.Println(*p1)  // 0 (zero value)
    *p1 = 42
    fmt.Println(*p1)  // 42
    
    // new() with structs
    type Person struct {
        Name string
        Age  int
    }
    
    p2 := new(Person)
    fmt.Println(p2)        // Memory address
    fmt.Println(*p2)       // { 0} (zero value)
    p2.Name = "Alice"      // Automatic dereferencing
    fmt.Println(p2.Name)   // Alice
    
    // make() - for slices, maps, channels only
    slice := make([]int, 5)     // len=5, cap=5
    slice2 := make([]int, 5, 10) // len=5, cap=10
    
    m := make(map[string]int)
    m["key"] = 42
    
    ch := make(chan int)
    
    // Equivalent alternatives
    var slice3 []int = make([]int, 5)
    var slice4 = make([]int, 5)
    slice5 := make([]int, 5)
    
    // new vs make comparison
    var p3 *[]int = new([]int)    // *p3 is nil slice
    var p4 []int = make([]int, 0) // p4 is empty slice
    
    fmt.Println(p3)      // Memory address
    fmt.Println(*p3)     // [] (nil slice)
    fmt.Println(p4)      // [] (empty slice)
    fmt.Println(*p3 == nil) // true
    fmt.Println(p4 == nil)  // false
}

// Memory management best practices
func memoryBestPractices() {
    // 1. Prefer stack allocation (automatic)
    func stackAllocation() {
        x := 42  // Allocated on stack
        // Automatically freed when function returns
    }
    
    // 2. Use make() for reference types
    func referenceTypes() {
        slice := make([]int, 0, 100)  // Pre-allocate capacity
        m := make(map[string]int, 100) // Pre-allocate capacity
    }
    
    // 3. Avoid unnecessary heap allocations
    func heapAllocation() {
        // This forces heap allocation
        x := 42
        return &x  // x escapes to heap
    }
    
    // 4. Reuse objects when possible
    func objectReuse() {
        // Use sync.Pool for expensive objects
        var bufferPool = sync.Pool{
            New: func() interface{} {
                return make([]byte, 0, 1024)
            },
        }
        
        // Get from pool
        buf := bufferPool.Get().([]byte)
        buf = buf[:0]  // Reset length
        
        // Use buffer...
        
        // Return to pool
        bufferPool.Put(buf)
    }
}
```

## Practice Projects

### Project 1: CLI Calculator

Let's implement a comprehensive CLI calculator:

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Calculator struct {
    history []string
    memory  float64
}

func NewCalculator() *Calculator {
    return &Calculator{
        history: make([]string, 0),
        memory:  0.0,
    }
}

func (c *Calculator) Add(a, b float64) float64 {
    result := a + b
    c.addToHistory(fmt.Sprintf("%.2f + %.2f = %.2f", a, b, result))
    return result
}

func (c *Calculator) Subtract(a, b float64) float64 {
    result := a - b
    c.addToHistory(fmt.Sprintf("%.2f - %.2f = %.2f", a, b, result))
    return result
}

func (c *Calculator) Multiply(a, b float64) float64 {
    result := a * b
    c.addToHistory(fmt.Sprintf("%.2f * %.2f = %.2f", a, b, result))
    return result
}

func (c *Calculator) Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    result := a / b
    c.addToHistory(fmt.Sprintf("%.2f / %.2f = %.2f", a, b, result))
    return result, nil
}

func (c *Calculator) Power(base, exponent float64) float64 {
    result := 1.0
    for i := 0; i < int(exponent); i++ {
        result *= base
    }
    c.addToHistory(fmt.Sprintf("%.2f ^ %.2f = %.2f", base, exponent, result))
    return result
}

func (c *Calculator) StoreMemory(value float64) {
    c.memory = value
    c.addToHistory(fmt.Sprintf("Stored %.2f in memory", value))
}

func (c *Calculator) RecallMemory() float64 {
    c.addToHistory(fmt.Sprintf("Recalled %.2f from memory", c.memory))
    return c.memory
}

func (c *Calculator) ClearMemory() {
    c.memory = 0
    c.addToHistory("Memory cleared")
}

func (c *Calculator) ShowHistory() {
    if len(c.history) == 0 {
        fmt.Println("No history available")
        return
    }
    fmt.Println("Calculation History:")
    for i, entry := range c.history {
        fmt.Printf("%d. %s\n", i+1, entry)
    }
}

func (c *Calculator) ClearHistory() {
    c.history = c.history[:0]
    fmt.Println("History cleared")
}

func (c *Calculator) addToHistory(entry string) {
    c.history = append(c.history, entry)
    if len(c.history) > 100 {  // Limit history to 100 entries
        c.history = c.history[1:]
    }
}

func parseFloat(s string) (float64, error) {
    return strconv.ParseFloat(strings.TrimSpace(s), 64)
}

func main() {
    calc := NewCalculator()
    scanner := bufio.NewScanner(os.Stdin)
    
    fmt.Println("Go Calculator")
    fmt.Println("Commands: add, sub, mul, div, pow, mem, recall, clear, history, quit")
    fmt.Println()
    
    for {
        fmt.Print("calc> ")
        if !scanner.Scan() {
            break
        }
        
        input := strings.TrimSpace(scanner.Text())
        if input == "" {
            continue
        }
        
        parts := strings.Fields(input)
        if len(parts) == 0 {
            continue
        }
        
        command := strings.ToLower(parts[0])
        
        switch command {
        case "quit", "exit", "q":
            fmt.Println("Goodbye!")
            return
            
        case "add":
            if len(parts) != 3 {
                fmt.Println("Usage: add <num1> <num2>")
                continue
            }
            a, err1 := parseFloat(parts[1])
            b, err2 := parseFloat(parts[2])
            if err1 != nil || err2 != nil {
                fmt.Println("Invalid numbers")
                continue
            }
            result := calc.Add(a, b)
            fmt.Printf("Result: %.2f\n", result)
            
        case "sub":
            if len(parts) != 3 {
                fmt.Println("Usage: sub <num1> <num2>")
                continue
            }
            a, err1 := parseFloat(parts[1])
            b, err2 := parseFloat(parts[2])
            if err1 != nil || err2 != nil {
                fmt.Println("Invalid numbers")
                continue
            }
            result := calc.Subtract(a, b)
            fmt.Printf("Result: %.2f\n", result)
            
        case "mul":
            if len(parts) != 3 {
                fmt.Println("Usage: mul <num1> <num2>")
                continue
            }
            a, err1 := parseFloat(parts[1])
            b, err2 := parseFloat(parts[2])
            if err1 != nil || err2 != nil {
                fmt.Println("Invalid numbers")
                continue
            }
            result := calc.Multiply(a, b)
            fmt.Printf("Result: %.2f\n", result)
            
        case "div":
            if len(parts) != 3 {
                fmt.Println("Usage: div <num1> <num2>")
                continue
            }
            a, err1 := parseFloat(parts[1])
            b, err2 := parseFloat(parts[2])
            if err1 != nil || err2 != nil {
                fmt.Println("Invalid numbers")
                continue
            }
            result, err := calc.Divide(a, b)
            if err != nil {
                fmt.Printf("Error: %v\n", err)
                continue
            }
            fmt.Printf("Result: %.2f\n", result)
            
        case "pow":
            if len(parts) != 3 {
                fmt.Println("Usage: pow <base> <exponent>")
                continue
            }
            base, err1 := parseFloat(parts[1])
            exp, err2 := parseFloat(parts[2])
            if err1 != nil || err2 != nil {
                fmt.Println("Invalid numbers")
                continue
            }
            result := calc.Power(base, exp)
            fmt.Printf("Result: %.2f\n", result)
            
        case "mem":
            if len(parts) != 2 {
                fmt.Println("Usage: mem <number>")
                continue
            }
            value, err := parseFloat(parts[1])
            if err != nil {
                fmt.Println("Invalid number")
                continue
            }
            calc.StoreMemory(value)
            
        case "recall":
            value := calc.RecallMemory()
            fmt.Printf("Memory: %.2f\n", value)
            
        case "clear":
            calc.ClearMemory()
            calc.ClearHistory()
            
        case "history":
            calc.ShowHistory()
            
        default:
            fmt.Println("Unknown command. Available commands:")
            fmt.Println("add, sub, mul, div, pow, mem, recall, clear, history, quit")
        }
    }
}
```

### Project 2: Contact Management System

```go
package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
    "strconv"
    "strings"
    "time"
)

type Contact struct {
    ID          int                    `json:"id"`
    Name        string                 `json:"name"`
    Email       string                 `json:"email"`
    Phone       string                 `json:"phone"`
    Address     Address                `json:"address"`
    Tags        []string               `json:"tags"`
    Notes       string                 `json:"notes"`
    CreatedAt   time.Time              `json:"created_at"`
    UpdatedAt   time.Time              `json:"updated_at"`
    CustomFields map[string]interface{} `json:"custom_fields"`
}

type Address struct {
    Street  string `json:"street"`
    City    string `json:"city"`
    State   string `json:"state"`
    Country string `json:"country"`
    ZIP     string `json:"zip"`
}

type ContactManager struct {
    contacts map[int]*Contact
    nextID   int
    filename string
}

func NewContactManager(filename string) *ContactManager {
    cm := &ContactManager{
        contacts: make(map[int]*Contact),
        nextID:   1,
        filename: filename,
    }
    cm.loadFromFile()
    return cm
}

func (cm *ContactManager) AddContact(name, email, phone string) *Contact {
    contact := &Contact{
        ID:           cm.nextID,
        Name:         name,
        Email:        email,
        Phone:        phone,
        Tags:         make([]string, 0),
        CreatedAt:    time.Now(),
        UpdatedAt:    time.Now(),
        CustomFields: make(map[string]interface{}),
    }
    
    cm.contacts[cm.nextID] = contact
    cm.nextID++
    cm.saveToFile()
    return contact
}

func (cm *ContactManager) GetContact(id int) (*Contact, bool) {
    contact, exists := cm.contacts[id]
    return contact, exists
}

func (cm *ContactManager) UpdateContact(id int, updates map[string]interface{}) error {
    contact, exists := cm.contacts[id]
    if !exists {
        return fmt.Errorf("contact with ID %d not found", id)
    }
    
    for field, value := range updates {
        switch field {
        case "name":
            if str, ok := value.(string); ok {
                contact.Name = str
            }
        case "email":
            if str, ok := value.(string); ok {
                contact.Email = str
            }
        case "phone":
            if str, ok := value.(string); ok {
                contact.Phone = str
            }
        case "notes":
            if str, ok := value.(string); ok {
                contact.Notes = str
            }
        }
    }
    
    contact.UpdatedAt = time.Now()
    cm.saveToFile()
    return nil
}

func (cm *ContactManager) DeleteContact(id int) error {
    if _, exists := cm.contacts[id]; !exists {
        return fmt.Errorf("contact with ID %d not found", id)
    }
    
    delete(cm.contacts, id)
    cm.saveToFile()
    return nil
}

func (cm *ContactManager) SearchContacts(query string) []*Contact {
    var results []*Contact
    query = strings.ToLower(query)
    
    for _, contact := range cm.contacts {
        if strings.Contains(strings.ToLower(contact.Name), query) ||
           strings.Contains(strings.ToLower(contact.Email), query) ||
           strings.Contains(strings.ToLower(contact.Phone), query) {
            results = append(results, contact)
        }
    }
    
    return results
}

func (cm *ContactManager) ListContacts() []*Contact {
    var contacts []*Contact
    for _, contact := range cm.contacts {
        contacts = append(contacts, contact)
    }
    return contacts
}

func (cm *ContactManager) AddTag(id int, tag string) error {
    contact, exists := cm.contacts[id]
    if !exists {
        return fmt.Errorf("contact with ID %d not found", id)
    }
    
    // Check if tag already exists
    for _, existingTag := range contact.Tags {
        if existingTag == tag {
            return nil // Tag already exists
        }
    }
    
    contact.Tags = append(contact.Tags, tag)
    contact.UpdatedAt = time.Now()
    cm.saveToFile()
    return nil
}

func (cm *ContactManager) RemoveTag(id int, tag string) error {
    contact, exists := cm.contacts[id]
    if !exists {
        return fmt.Errorf("contact with ID %d not found", id)
    }
    
    for i, existingTag := range contact.Tags {
        if existingTag == tag {
            contact.Tags = append(contact.Tags[:i], contact.Tags[i+1:]...)
            contact.UpdatedAt = time.Now()
            cm.saveToFile()
            return nil
        }
    }
    
    return fmt.Errorf("tag '%s' not found", tag)
}

func (cm *ContactManager) saveToFile() error {
    file, err := os.Create(cm.filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    return encoder.Encode(cm.contacts)
}

func (cm *ContactManager) loadFromFile() error {
    file, err := os.Open(cm.filename)
    if err != nil {
        if os.IsNotExist(err) {
            return nil // File doesn't exist yet
        }
        return err
    }
    defer file.Close()
    
    decoder := json.NewDecoder(file)
    contacts := make(map[int]*Contact)
    
    if err := decoder.Decode(&contacts); err != nil {
        return err
    }
    
    cm.contacts = contacts
    
    // Find next ID
    maxID := 0
    for id := range contacts {
        if id > maxID {
            maxID = id
        }
    }
    cm.nextID = maxID + 1
    
    return nil
}

func printContact(contact *Contact) {
    fmt.Printf("ID: %d\n", contact.ID)
    fmt.Printf("Name: %s\n", contact.Name)
    fmt.Printf("Email: %s\n", contact.Email)
    fmt.Printf("Phone: %s\n", contact.Phone)
    if contact.Address.Street != "" {
        fmt.Printf("Address: %s, %s, %s %s, %s\n",
            contact.Address.Street,
            contact.Address.City,
            contact.Address.State,
            contact.Address.ZIP,
            contact.Address.Country)
    }
    if len(contact.Tags) > 0 {
        fmt.Printf("Tags: %s\n", strings.Join(contact.Tags, ", "))
    }
    if contact.Notes != "" {
        fmt.Printf("Notes: %s\n", contact.Notes)
    }
    fmt.Printf("Created: %s\n", contact.CreatedAt.Format("2006-01-02 15:04:05"))
    fmt.Printf("Updated: %s\n", contact.UpdatedAt.Format("2006-01-02 15:04:05"))
    fmt.Println(strings.Repeat("-", 40))
}

func main() {
    cm := NewContactManager("contacts.json")
    scanner := bufio.NewScanner(os.Stdin)
    
    fmt.Println("Contact Management System")
    fmt.Println("Commands: add, list, search, update, delete, tag, untag, quit")
    fmt.Println()
    
    for {
        fmt.Print("contacts> ")
        if !scanner.Scan() {
            break
        }
        
        input := strings.TrimSpace(scanner.Text())
        if input == "" {
            continue
        }
        
        parts := strings.Fields(input)
        if len(parts) == 0 {
            continue
        }
        
        command := strings.ToLower(parts[0])
        
        switch command {
        case "quit", "exit", "q":
            fmt.Println("Goodbye!")
            return
            
        case "add":
            fmt.Print("Name: ")
            scanner.Scan()
            name := strings.TrimSpace(scanner.Text())
            
            fmt.Print("Email: ")
            scanner.Scan()
            email := strings.TrimSpace(scanner.Text())
            
            fmt.Print("Phone: ")
            scanner.Scan()
            phone := strings.TrimSpace(scanner.Text())
            
            contact := cm.AddContact(name, email, phone)
            fmt.Printf("Contact added with ID: %d\n", contact.ID)
            
        case "list":
            contacts := cm.ListContacts()
            if len(contacts) == 0 {
                fmt.Println("No contacts found")
                continue
            }
            
            fmt.Printf("Found %d contacts:\n\n", len(contacts))
            for _, contact := range contacts {
                printContact(contact)
            }