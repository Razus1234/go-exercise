package main

import "fmt"

func exercise7_method(persons Person) {

	fmt.Println("✅ Exercise 7: เพิ่ม method ให้ struct Person")
	fmt.Println("Greeting:", persons.Greet())
}

func (p Person) Greet() string {
	return "Hi " + p.Name
}
