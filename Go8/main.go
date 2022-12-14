package main

import (
	"fmt"
	"log"
	"net/http"
)

type messageHandler struct {
	message string
}

func main() {

	mux := http.NewServeMux()

	x1 := &messageHandler{"Bu bir mesaj!"}
	x2 := &messageHandler{"Bu da bir mesaj!"}

	mux.Handle("/bir", x1)
	mux.Handle("/iki", x2)

	log.Println("Listening...")

	http.ListenAndServe(":8080", mux)

}

func (x *messageHandler) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	fmt.Fprint(res, x.message)
}
