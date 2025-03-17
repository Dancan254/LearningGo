package main

import "fmt"

func main() {
	sum := sumOfNums()
	fmt.Println(sum)

	fmt.Println(sum2())
}
func sumOfNums() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}
func sum2() int {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	return sum
}

// for is Go's while loop
func sum3() int {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	return sum
}
