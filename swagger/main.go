package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe(":8080", nil)
}
