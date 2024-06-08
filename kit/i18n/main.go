package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	err := NewUserNotFoundErr(123)
	// 某些业务处零零落落、
	//err, _ := someFunc()
	fmt.Fprintf(w, FormatErr(err))
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
