package main

import "fmt"

func main() {
	fmt.Println("Problem 1")
	problem1()

	fmt.Println("Problem 2")
	probelm2()

	fmt.Println("Problem 3")
	problem3()
}

// Write a program that declares an integer variable called i with the value 20.
// Assign i to a floating-point variable named f.
// Print out i and f.
func problem1() {
	i := 20
	var f float32 = float32(i)
	fmt.Println(i)
	fmt.Println(f)
}

// Write a program that declares a constant called value that can be assigned to both an integer and a floating-point variable.
// Assign it to an integer called i and a floating-point variable called f.
// Print out i and f.
func probelm2() {
	const value = 10
	var i int64 = value
	var f float64 = value
	fmt.Println(i)
	fmt.Println(f)
}

// Write a program with three variables, one named b of type byte, one named smallI of type int32, and one named bigI of type uint64.
// Assign each variable the maximum legal value for its type; then add 1 to each variable.
// Print out their values.
func problem3() {
	var b byte = byte(255)
	var smallI int32 = 2147483647
	var bigI int64 = 9223372036854775807

	b++
	smallI++
	bigI++

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}
