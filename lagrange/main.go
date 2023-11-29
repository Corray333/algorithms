package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	fmt.Println("Server is listening 3001...")
	http.ListenAndServe("127.0.0.1:3001", nil)
}
