package main

import (
	// "encoding/json"
	// "fmt"
	// "log"
	// "net/http"

	// "github.com/gorilla/mux"

	"log"
	"net/http"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })

	// log.Fatal(http.ListenAndServe(":8080", nil))

	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", Index)
	// router.HandleFunc("/todos", TodoIndex)
	// router.HandleFunc("/todos/{todoId}", TodoShow)

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
