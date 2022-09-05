package main

import (
	"fmt"
	"net/http"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))

	x := r.URL.Path[1:]
	fmt.Print(x)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("İndex Page"))
}
func main() {
	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8080", nil)
}
