package main

import "fmt"

func main() {
	personOne := MakerPerson("Sean", "Blizard", 21)
	fmt.Println(personOne)

	personTwo := MakerPersonPointer("Kelly", "Blizard", 48)
	fmt.Println(personTwo)

	stringSlice := []string{"Hello", "world!"}
	fmt.Println("stringSlice:", stringSlice)
	UpdateSlice(stringSlice, "Sup!")
	fmt.Println("stringSlice:", stringSlice)
	GrowSlice(stringSlice, "Howdy!")
	fmt.Println("stringSlice:", stringSlice)
}

//“Create a struct named Person with three fields: FirstName and LastName of type string and Age of type int.
// Write a function called MakePerson that takes in firstName, lastName, and age and returns a Person.
// Write a second function MakePersonPointer that takes in firstName, lastName, and age and returns a *Person.
// Call both from main. Compile your program with go build -gcflags="-m".
// This both compiles your code and prints out which values escape to the heap.
// Are you surprised about what escapes?

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakerPerson(firstName string, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakerPersonPointer(firstName string, lastName string, age int) *Person {
	person := MakerPerson(firstName, lastName, age)
	return &person
}

// Write two functions. The UpdateSlice function takes in a []string and a string.
// It sets the last position in the passed-in slice to the passed-in string.
// At the end of UpdateSlice, print the slice after making the change.
// The GrowSlice function also takes in a []string and a string.
// It appends the string onto the slice. At the end of GrowSlice, print the slice after making the change.
// Call these functions from main. Print out the slice before each function is called and after each function is called.
// Do you understand why some changes are visible in main and why some changes are not?
func UpdateSlice(stringSlice []string, str string) {
	stringSlice[len(stringSlice)-1] = str
	fmt.Println("Updated slice:", stringSlice)
}

func GrowSlice(stringSlice []string, str string) {
	stringSlice = append(stringSlice, str)
	fmt.Println("Grown slice: ", stringSlice)
}
