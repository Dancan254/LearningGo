package main

import (
	"fmt"
	"math"
)

func main() {

	age := 18
	if age >= 23 {
		fmt.Println("You can vote")
		fmt.Println("You can drive")
	} else if age >= 18 && age < 23 {
		fmt.Println("You can't drive")
		fmt.Println("You can vote")
	} else {
		fmt.Println("You can't vote")
		fmt.Println("You can't drive")
	}

	//using short statement
	if n := 10; n < 20 {
		fmt.Println(n)
	}

	if num := 9; num%2 == 0 {
		fmt.Printf("%d is even", num)
	}

	messageLen := 10
	maxMessageLen := 20
	fmt.Printf("Trimh to send a message of length %d", messageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}

	fmt.Println(pow(2, 3, 6))
	fmt.Println(power(3, 2, 10))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func power(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here
	return lim
}
