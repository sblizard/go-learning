package main

import (
	"fmt"
	"math/rand"
)

func main() {
	exercise1()
	exercise2()
}

// Write a for loop that puts 100 random numbers between 0 and 100 into an int slice.
// Loop over the slice you created in exercise 1. For each value in the slice, apply the following rules:
// If the value is divisible by 2, print “Two!”
// If the value is divisible by 3, print “Three!”
// IIf the value is divisible by 2 and 3, print “Six!”. Don’t print anything else.
// Otherwise, print “Never mind”.
func exercise1() {
	randInts := make([]int, 100)
	for i := range 100 {
		randInts[i] = rand.Intn(101)
	}
	for _, v := range randInts {
		if v%2 == 0 && v%3 == 0 {
			fmt.Println("Six!")
			continue
		} else if v%2 == 0 {
			fmt.Println("Two!")
			continue
		} else if v%3 == 0 {
			fmt.Println("Three!")
			continue
		}
		fmt.Println("Never mind.")
	}
}

// Start a new program. In main, declare an int variable called total.
// Write a for loop that uses a variable named i to iterate from 0 (inclusive) to 10 (exclusive).
// The body of the for loop should be as follows:
// total := total + i
// fmt.Println(total)
func exercise2() {
	var total int
	for i := range 11 {
		total := total + i
		fmt.Println(total)
	}
	fmt.Println("Bad program, it is shadowing total.")
}
