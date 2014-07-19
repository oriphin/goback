package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"net/http"
)

func main() {
	go StartServer()
	fmt.Scanln()
}

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Go Back")
	})

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":4000")
}
