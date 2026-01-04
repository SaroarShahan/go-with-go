package main

import "fmt"

func printGreetings() {
	fmt.Println("Hello! Welcome to our application.")
}

func getUserName() string {
	var name string

	fmt.Println("Please enter your name: ")
	fmt.Scanln(&name)

	return name
}

func getTwoNumbers() (int, int) {
	var num1, num2 int

	fmt.Println("Enter two numbers: ")
	fmt.Scanln(&num1, &num2)

	return num1, num2
}

func getSum(a int, b int) int {
	return a + b
}

func printResults(name string, num1 int, num2 int, sum int) {
	fmt.Printf("Hello, %s! The sum of %d and %d is %d.\n", name, num1, num2, sum)
}

func printFarewell(name string) {
	fmt.Printf("Thank you %s for using the application!\n", name)
	fmt.Println("Goodbye!")
}

func main() {
	// Welcome message
	printGreetings()

	// Get user name
	name := getUserName()

	// Get two numbers
	num1, num2 := getTwoNumbers()

	// Calculate sum
	result := getSum(num1, num2)

	// Display results
	printResults(name, num1, num2, result)
	
	// Farewell message
	printFarewell(name)
}