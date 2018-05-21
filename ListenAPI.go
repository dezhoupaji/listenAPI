package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/showA", showA)
	http.HandleFunc("/showB", showB)
	err := http.ListenAndServe(":40000", nil)
	if err != nil {
		println(err)
	}
}

func showA(w http.ResponseWriter, r *http.Request) {
	println("receive request A")
}

func showB(w http.ResponseWriter, r *http.Request) {
	println("receive request B")
}
