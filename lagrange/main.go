package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.ListenAndServe("127.0.0.1:3000", nil)
}
