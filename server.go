package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	go StartServer()
	fmt.Scanln()
}

func StartServer() {

	router := mux.NewRouter()
	router.HandleFunc("/api", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Go Back")
	})

	static := negroni.NewStatic(http.Dir("public"))
	static.IndexFile = "templates/index.html"

	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		static)
	n.UseHandler(router)

	n.Run(":4000")
}
