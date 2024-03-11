package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)


func main(){
	fmt.Println("Server is running on port 8080")

	h1 := func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello, World!")
		io.WriteString(w, r.Method)

	}
	http.HandleFunc("/",h1)

	log.Fatal(http.ListenAndServe(":8080",nil))
}