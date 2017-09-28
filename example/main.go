package main

import (
	"fmt"
	"net/http"

	"github.com/kidandcat/armony"
)

func reparse(w *http.ResponseWriter, r *http.Request, ss *armony.Session) (string, interface{}) {
	fmt.Println("reparsing templates")

	armony.ParseTemplates([]string{
		"views",
		"components",
	})

	ss.Set("username", "World")

	return "", nil
}

func index(w *http.ResponseWriter, r *http.Request, ss *armony.Session) (string, interface{}) {

	user := ss.Get("username")
	return "template:index", armony.Data{
		"Username": user,
	}
}

func main() {
	armony.LoadDatabase()
	armony.ParseTemplates([]string{
		"views",
		"components",
	})
	armony.AddRoute("/reloadTemplates", reparse)
	armony.AddRoute("/", index)

	http.HandleFunc("/", armony.Handler)

	server := http.Server{
		Addr: "127.0.0.1:80",
	}
	fmt.Println("Server listening in", "127.0.0.1:80")
	server.ListenAndServe()
}
