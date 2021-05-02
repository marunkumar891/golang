package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hey there your requested url is : %s", r.URL.Path)
	// this fprintf function is specially used to format and writes in W
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", index)
	http.ListenAndServe(":8080", mux)
}
