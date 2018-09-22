package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	AppName string
}

func main() {
	app := App{AppName: "Hamburgers 'R' Us"}
	r := mux.NewRouter()
	r.HandleFunc("/", app.handler).Methods(http.MethodGet)
	http.ListenAndServe(":8080", r)
}

func (app *App) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! - %s", app.AppName)
}
