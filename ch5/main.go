package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria

	calc()
}

// The simple calculator program doesn’t handle one error case: division by zero.
// Change the function signature for the math operations to return both an int and an error.
// In the div function, if the divisor is 0, return errors.New("division by zero") for the error.
// In all other cases, return nil. Adjust the main function to check for this error.
func add(i int, j int) (int, error) { return i + j, nil }

func sub(i int, j int) (int, error) { return i - j, nil }

func mul(i int, j int) (int, error) { return i * j, nil }

func div(i int, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("division by zero")
	}
	return i / j, nil
}

var opMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func calc() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
	}
}

// Write a function called fileLen that has an input parameter of type string and returns an int and an error.
// The function takes in a filename and returns the number of bytes in the file.
// If there is an error reading the file, return the error. Use defer to make sure the file is closed properly.
func fileLen(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	data := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		data = append(data, buf[:n]...)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return 0, err
		}
	}
	return len(data), nil
}

// Write a function called prefixer that has an input parameter of type string and returns a function that has an input
// parameter of type string and returns a string.
// The returned function should prefix its input with the string passed into prefixer.
// Use the following main function to test prefixer:
func prefixer(prefix string) func(string) string {
	return func(s string) string {
		return prefix + " " + s
	}
}
