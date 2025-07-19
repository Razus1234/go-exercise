package main

import "fmt"

func exercise8_crud(persons []Person) []Person {
	fmt.Println("✅ Exercise 8: สร้าง CRUD ใน slice")
	fmt.Println("Create, Read, Update, Delete operations on slice of persons")
	// Create
	persons = AddPerson(persons, Person{Name: "David", Age: 28})
	// Read
	fmt.Println("Find person by name 'Alice':", FindPersonByName(persons, "Alice"))
	// Update
	UpdateAge(persons, "Bob", 32)
	// Delete
	persons = DeletePersonByName(persons, "Charlie")
	fmt.Println("Persons after CRUD operations:")
	for i, person := range persons {
		fmt.Printf("No: %d Name: %s, Age: %d\n", i+1, person.Name, person.Age)
	}
	return persons
}

func AddPerson(slice []Person, p Person) []Person {
	// Append the new person to the slice
	slice = append(slice, p)
	fmt.Println("Added:", p.Name)
	return slice
}

func FindPersonByName(persons []Person, name string) Person {
	// Loop through the slice to find the person by name
	for _, person := range persons {
		if person.Name == name {
			return person
		}
	}
	return Person{} // Return an empty Person if not found
}

func UpdateAge(persons []Person, name string, newAge int) {
	// Loop through the slice to find the person by name and update their age
	for i, person := range persons {
		if person.Name == name {
			persons[i].Age = newAge
			fmt.Printf("Updated %s's age to %d\n", name, newAge)
			return
		}
	}
	fmt.Println("Person not found:", name)
}

func DeletePersonByName(persons []Person, name string) []Person {
	// Loop through the slice to find the person by name and delete them
	for i, person := range persons {
		if person.Name == name {
			persons = append(persons[:i], persons[i+1:]...)
			fmt.Println("Deleted:", name)
			return persons
		}
	}
	fmt.Println("Person not found:", name)
	return persons
}
