package main

import (
	"fmt"
	"sync"
)

func exercise9_goroutine() {
	var wg sync.WaitGroup // คือ WaitGroup เพื่อรอให้ goroutine เสร็จสิ้น

	fmt.Println("✅ Exercise 9: สร้าง Goroutine")
	wg.Add(1) // เพิ่มจำนวน goroutine ที่ต้องรอ

	// สร้าง goroutine ใหม่
	go func() {
		defer wg.Done()
		fmt.Println("Hello from goroutine!")
	}()

	wg.Wait() // รอให้ goroutine เสร็จสิ้น
	// เมื่อ goroutine เสร็จสิ้น จะทำงานต่อไปที่นี่
	fmt.Println("Goroutine finished execution.")
}
