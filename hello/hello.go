package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"Hello,World!")
}

func main() {
	http.HandleFunc("/hello",helloHandler)
	fmt.Println("Server starting on port :8800...")
	http.ListenAndServe(":8800",nil)
}