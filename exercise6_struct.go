package main

import "fmt"

func exercise6_struct() []Person {
	fmt.Println("✅ Exercise 6: สร้าง struct Person")
	fmt.Println("ใส่ไว้ใน slice แล้ว loop แสดงผล")
	persons := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}

	fmt.Println("Persons:")
	// Loop through the slice of persons and print their details
	// Note: The index is not used in the print statement, so we can use an underscore

	_ = 0
	for i, person := range persons {
		// Using the index i to print the person's number
		fmt.Printf("No: %d Name: %s, Age: %d\n", i+1, person.Name, person.Age)
	}
	return persons
}

type Person struct {
	Name string
	Age  int
}
