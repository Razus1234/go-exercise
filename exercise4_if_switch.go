package main

import "fmt"

func exercise4_if_switch() {

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
}
