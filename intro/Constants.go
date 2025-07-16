package main

import "fmt"

func main() {
	// constants are compile-time values with unique features
	//Basic Constants
	const Pi = 3.14
	const Name = "GoLang"
	const Active = true

	//typed constants
	const TypedInt int = 42
	const TypedFloat float64 = 3.14

	//untyped constants - these can be used in any context where a value of the appropriate type is expected
	// they are flexible and can be used in various operations without explicit type conversion
	const UntypedInt = 100       // can be used as int, int32, int64, etc.
	const UntypedFloat = 2.71828 // can be used as float32, float64, etc.

	// Constants expressions (evaluated at compile time)
	const (
		SecondsPerMinute = 60
		MinutesPerHour   = 60
		SecondPerHour    = SecondsPerMinute * MinutesPerHour
	)

	// iota - a predeclared identifier that represents successive untyped integer constants
	//basically it is a counter that increments by 1 for each constant declaration in a const block
	const (
		Sunday = iota // 0
		Monday        // 1
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	//Advanced iota patterns
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
	)
	// Example usage of constants
	fmt.Println("Pi:", Pi)
	zeroValues()

}
func zeroValues() {
	var i int             // 0
	var f float64         // 0.0
	var b bool            // false
	var s string          // ""
	var ptr *int          // nil
	var slice []int       // nil
	var m map[string]int  // nil
	var ch chan int       // nil
	var fn func()         // nil
	var iface interface{} // nil

	// Structs get zero values for all fields
	type Person struct {
		Name string
		Age  int
	}
	var p Person // {Name: "", Age: 0}

	// Arrays get zero values for all elements
	var arr [5]int // [0, 0, 0, 0, 0]

	fmt.Println("Zero values:")
	fmt.Println("int:", i)
	fmt.Println("float64:", f)
	fmt.Println("bool:", b)
	fmt.Println("string:", s)
	fmt.Println("pointer:", ptr)
	fmt.Println("slice:", slice)
	fmt.Println("map:", m)
	fmt.Println("channel:", ch)
	fmt.Println("function:", fn)
	fmt.Println("interface:", iface)
	fmt.Println("struct:", p)
	fmt.Println("array:", arr)
}
