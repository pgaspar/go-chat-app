// Go is procedural. There are no classes, only structs

// Install fresh (go get github.com/pilu/fresh), cd to work dir & run `fresh`

// := does type inference
// = does not. a := 1; a = "abc" will break.
// var i float32 declares i as a float
// var i = 1 also infers

// Struct variables default to "" (if they're strings, I suppose)
// type Index struct {
//  Name string // default: ""
//}

// If a variable starts with capital letter it's public.
// If it doesn't, it's private, so you can't use it in the templates.

// This is weird. If you include a private var in a template it stops execution and gives you half a page, basically - simply stopping execution.
// However, no proper error is presented...

// gorillatoolkit.org - series of packages you can use to build web apps
// gorilla/mux -> routing (go get github.com/gorilla/mux)

package main

import "net/http"
import "html/template"
import "github.com/gorilla/mux"

var templates = template.Must(template.ParseGlob("templates/*.tmpl"))

type User struct {
  Username string
}

// array with User structs, initialized with a single User
// with "admin" as username
var users = []User { User{"admin"} }

func handler(w http.ResponseWriter, r *http.Request) {
  type Header struct {
    Title string
  }
  type Index struct {
    Name string
    Header Header
  }

  data := Index {
    Name: "pedro",
    Header: Header {
      Title: "yoooo goooo!",
    },
  }

  templates.ExecuteTemplate(w, "index", &data)
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", handler)
  r.HandleFunc("/users", handler).Methods("GET")
  http.ListenAndServe(":8080", r)
}
