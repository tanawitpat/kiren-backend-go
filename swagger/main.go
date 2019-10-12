package main

import (
	"fmt"
	"net/http"
)

// Load the api_spec.yaml from <project_root>/swagger/dist/api_spec.yaml
// and preview it inside the browser.

func main() {
	fs := http.FileServer(http.Dir("swagger/dist"))
	http.Handle("/", http.StripPrefix("/", fs))
	fmt.Println("The api_spec is being served on http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
