package main

import (
	"fmt"
	"math"
)

// when you declare sth you must use otherwise you get a compile time error
func main() {
	//variables
	var a int = 10 //default value is 0
	//multiple declaration, must be of same type though
	var b, c int = 1, 2
	fmt.Println(b, c)
	fmt.Println(a)

	//inlining different types
	var (
		aa int    = 10
		bb int    = 20
		cc string = "hello"
	)
	fmt.Println(aa, bb, cc)

	/*
			//type inference -> determining the data type
				basics types include:

				Integers			int
				Floats				float64
				Complex Numbers		complex128
				Strings				string
				Booleans			bool
				Characters			int32 or rune

		For other types such as Array, Pointer, Structure, etc, type Inference will happen based on the value.
	*/

	var t = 123                  //infers to an int as the value is an integer
	var u = "circle"             //string
	var v = 5.6                  //float64
	var w = true                 //bool
	var x = 'a'                  //rune
	var y = 3 + 5i               //complex128
	var z = sample{name: "test"} //type inferred will be a main.Sample

	fmt.Printf("Type: %T Value: %v\n\n", t, t)
	fmt.Printf("Type: %T Value: %v\n\n", u, u)
	fmt.Printf("Type: %T Value: %v\n\n", v, v)
	fmt.Printf("Type: %T Value: %v\n\n", w, w)
	fmt.Printf("Type: %T Value: %v\n\n", x, x)
	fmt.Printf("Type: %T Value: %v\n\n", y, y)
	fmt.Printf("Type: %T Value: %v\n\n", z, z)

	//short variable declaration := -> both var and type can be omitted

	num := 123
	circ := "circle"
	fl := 5.6
	boo := true
	cha := 'a'
	com := 3 + 5i
	sam := sample{name: "test"}
	fmt.Printf("Type: %T Value: %v\n\n", num, num)
	fmt.Printf("Type: %T Value: %v\n\n", circ, circ)
	fmt.Printf("Type: %T Value: %v\n\n", fl, fl)
	fmt.Printf("Type: %T Value: %v\n\n", boo, boo)
	fmt.Printf("Type: %T Value: %v\n\n", cha, cha)
	fmt.Printf("Type: %T Value: %v\n\n", com, com)
	fmt.Printf("Type: %T Value: %v\n\n", sam, sam)

	//Note: the operator is only available within a function
	// once a variable is declared using the operator it cannot be redeclared
	//can also be used to declare multiple variables
	//can be used again if a particular variable is new

	ab, cd := 1, 2
	cd, ef := 3, 4
	fmt.Println(ab, cd, ef)

	//KEY POINTS
	/*
		a variable declared in the inner scope having the same name as one in the outer scope will overshadow it
		variables can be assigned as expressions or function call
		A variable once intialized with a particular type, cannot be assigned a value of different type later.
	*/
	ans := math.Sqrt(100)
	fmt.Println(ans)

	/*
		If this variable name starts with a lowercase letter then it can be accessed from within the package which contains this variable definition.
		If the variable name starts with a uppercase letter then it can be accessed from outside different package other than which it is declared.
		local variable -> These variables only live till the end of the block or a function in which they are declared. After that, they are Garbage Collected.
	*/

	//constants

	const pi = 3.14

	fmt.Printf("pi = %v\n", pi)

	//computed constants -> computed at compile time they can be left unused i guess, they do not the unused error
	const firstname = "John"
	const lastname = "Doe"
	const fullname = firstname + " " + lastname
	fmt.Println(fullname)

	//interpolation
	/*
	  Interpolation in Go:
	  - Go does not support string interpolation directly like some other languages (e.g., Python, JavaScript).
	  - Instead, Go uses the `fmt` package to format strings.
	  - The `fmt.Sprintf` function is commonly used for creating formatted strings.
	  - Placeholders like `%v`, `%d`, `%s`, etc., are used within the format string to indicate where values should be inserted.

	*/
	name := "John"
	age := 30
	formattedString := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
	fmt.Println(formattedString)

	//interpolating different data types
	price := 100.50
	quantity := 2
	total := price * float64(quantity)
	formattedString = fmt.Sprintf("The total price of %d items is %f", quantity, total)
	fmt.Println(formattedString)

	//interpolation with booleans
	isMember := true
	formattedString = fmt.Sprintf("Membership status: %t", isMember)
	fmt.Println(formattedString)

}

type sample struct {
	name string
}
