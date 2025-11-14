package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
}

// Write a program that defines a variable named greetings of type slice of strings with the following values:
// "Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт".
// Create a subslice containing the first two values; a second subslice with the second, third, and fourth values; and a third subslice with the fourth and fifth values.
// Print out all four slices.
func exercise1() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	a := greetings[:2]
	b := greetings[1:4]
	c := greetings[3:]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// Write a program that defines a string variable called message with the value "Hi 🤬 and 👨" and prints the fourth rune in it as a character, not a number.
func exercise2() {
	message := "Hi 🤬 and 👨"
	runes := []rune(message)
	fmt.Println(string(runes[9]))
}

// Write a program that defines a struct called Employee with three fields: firstName, lastName, and id.
// The first two fields are of type string, and the last field (id) is of type int.
// Create three instances of this struct using whatever values you’d like.
// Initialize the first one using the struct literal style without names, the second using the struct literal style with names, and the third with a var declaration.
// Use dot notation to populate the fields in the third struct. Print out all three structs.
func exercise3() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	employee1 := Employee{
		"Sean",
		"Blizard",
		0,
	}

	employee2 := Employee{
		firstName: "Santa",
		lastName:  "Clause",
		id:        1,
	}

	var employee3 struct {
		firstName string
		lastName  string
		id        int
	}
	employee3.firstName = "Jonny"
	employee3.lastName = "Appleseed"
	employee3.id = 2

	fmt.Println(employee1)
	fmt.Println(employee2)
	fmt.Println(employee3)
}
