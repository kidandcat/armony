package armony

import (
	"net/http"
	"net/url"
)

type routes map[string]Controller

// Controller : A type representing a controller funcion
type Controller func(*http.ResponseWriter, *http.Request, *Session) string

// Routes : all routes
var Routes routes

var routesInitialized = false

// Handler : Armony router handler
func Handler(w http.ResponseWriter, r *http.Request) {
	ss := LoadSession(&w, r)

	u, _ := url.Parse(r.RequestURI)

	if fn, ok := Routes[u.EscapedPath()]; ok {
		res := fn(&w, r, &ss)

		if res != "" {
			err := Templates.ExecuteTemplate(w, res, nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

// AddRoute : Adds a new route
func AddRoute(path string, controller Controller) {
	if !routesInitialized {
		Routes = make(routes)
		routesInitialized = true
	}
	Routes[path] = controller
}

// RemoveRoute : Removes a new route
func RemoveRoute(path string, controller Controller) {
	delete(Routes, path)
}
