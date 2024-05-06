git ppackage main

import (
	"fmt"
	"net/http"
)

func main() {
	http.NewServeMux()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "todo")
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "wip")
	})

<<<<<<< HEAD
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
=======
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
>>>>>>> private-key
		fmt.Fprintf(w, "wip")
	})

	http.ListenAndServe(":8080", nil)
}
