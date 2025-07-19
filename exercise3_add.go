package main

import "fmt"

func exercise3_add() {
	fmt.Println("✅ Exercise 3: ฟังก์ชันบวกเลข")
	fmt.Println("The sum of 5 and 3 is:", add(5, 3))
}

func add(x, y int) int {
	return x + y
}
