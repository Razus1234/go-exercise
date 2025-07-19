package main

import "fmt"

func main() {

	exercise1_hello()
	exercise2_variable()
	exercise3_add()
	exercise4_if_switch()
	exercise5_looping()
	persons := exercise6_struct()
	exercise7_method(persons[0])
	persons = exercise8_crud(persons)

	fmt.Println("✅ All exercises completed successfully!")
}
