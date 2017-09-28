# armony 
[![GPLv3](https://img.shields.io/aur/license/yaourt.svg)](LICENSE) [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](README.md#contribute)

Golang Web Framework for fast development

### Install

```
go get "github.com/kidandcat/armony"
```

### Example
```
package main

import (
	"fmt"
	"net/http"

	"github.com/kidandcat/armony"
)

func reparse(w *http.ResponseWriter, r *http.Request, ss *armony.Session) (string, interface{}) {
	fmt.Println("reparsing templates")

  // Reload Templates
	armony.ParseTemplates([]string{
		"views",
		"components",
	})

  // Set session variable
	ss.Set("username", "World")

  // Return nothing
	return "", nil
}

func index(w *http.ResponseWriter, r *http.Request, ss *armony.Session) (string, interface{}) {
  // Get sesion variable
	user := ss.Get("username")
  
  // Return commands and data
  // Render template "index" with data 
	return "template:index", armony.Data{
		"Username": user,
	}
}

func main() {
  // Load database for sessions
	armony.LoadDatabase()
  // Parse folders for templates (templates must end in .html)
	armony.ParseTemplates([]string{
		"views",
		"components",
	})
  // Add routes with handlers
	armony.AddRoute("/reloadTemplates", reparse)
	armony.AddRoute("/", index)

  // Let the Armony Handler handle all routes (or the ones you want)
	http.HandleFunc("/", armony.Handler)

  // Proceed like always
	server := http.Server{
		Addr: "127.0.0.1:80",
	}
	fmt.Println("Server listening in", "127.0.0.1:80")
	server.ListenAndServe()
}

```

### Documentation

Soon

### Contribute

Please send a GitHub Pull Request to armony with a clear list of what you've done (read more about pull requests). Please make sure all of your commits are atomic (one feature per commit).

Always write a clear log message for your commits.
