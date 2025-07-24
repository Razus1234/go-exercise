package main

import (
	"fmt"
	"net/http"
)

func webserver() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	fmt.Println("Starting web server on :8080")

	//print the url if start successfully
	fmt.Println("Visit http://localhost:8080 in your browser")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	} else {
		fmt.Println("Web server started successfully!")
	}
}
