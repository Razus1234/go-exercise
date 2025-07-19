package main

import "fmt"

func main() {

	fmt.Println("✅ Exercise 1: แสดงคำทักทาย")
	fmt.Println("Hello, Golang!")

	fmt.Println("✅ Exercise 2: ประกาศและพิมพ์ตัวแปร")
	name := "John"
	age := 30
	fmt.Println("My name is:", name, "I am", age, "years old")

	fmt.Println("✅ Exercise 3: ฟังก์ชันบวกเลข")
	fmt.Println("The sum of 5 and 3 is:", add(5, 3))

	fmt.Println("✅ Exercise 4: ใช้เงื่อนไข if และ switch ")
	fmt.Println("รับคะแนน int แล้วพิมพ์เกรด A, B, C, F")
	fmt.Println("Using if-else:")
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	fmt.Println("Using switch case:")
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")

	}

	fmt.Println("✅ Exercise 5: Looping with for")
	for i := 1; i <= 5; i++ {
		fmt.Println("Loop iteration:", i)
	}

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

	fmt.Println("✅ Exercise 7: เพิ่ม method ให้ struct Person")
	fmt.Println("Greeting:", persons[0].Greet())

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
	fmt.Println("✅ All exercises completed successfully!")
}

func add(x, y int) int {
	return x + y
}

type Person struct {
	Name string
	Age  int
}

//✅ Exercise 7: เพิ่ม method ให้ struct
func (p Person) Greet() string {
	return "Hi " + p.Name
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
