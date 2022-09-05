package main

import (
	"io"
	"net/http"
)

type ironman int

type wolverine int

func main() {

	var i ironman
	var w wolverine
	mux := http.NewServeMux()

	mux.Handle("/ironman", i)
	mux.Handle("/wolverine", w)

	http.ListenAndServe(":8080", mux)

}

func (x ironman) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Mr.Iron!")
}

func (x wolverine) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Mr.Wolverine!")
}
