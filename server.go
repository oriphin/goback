package main

import (
	"encoding/json"
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
	router.HandleFunc("/api", TestAPIHandler)

	static := negroni.NewStatic(http.Dir("public"))
	static.IndexFile = "templates/index.html"

	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		static)
	n.UseHandler(router)

	n.Run(":4000")
}

type Human struct {
	Username string
}

func TestAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := &Human{Username: "Oriphin"}
	j := json.NewEncoder(w)
	j.Encode(user)
}
