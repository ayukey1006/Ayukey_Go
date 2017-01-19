package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", Hey)
	http.ListenAndServe(":8081", nil)
}

func Hey(w http.ResponseWriter, r *http.Request) {

}
