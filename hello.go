// Go is procedural. There are no classes, only structs

// Install fresh (go get github.com/pilu/fresh), cd to work dir & run `fresh`

// := does type inference
// = does not. a := 1; a = "abc" will break.
// var i float32 declares i as a float
// var i = 1 also infers

// This is weird. If you include a private var in a template it stops execution and gives you half a page, basically - simply stopping execution.
// However, no proper error is presented...

package main

import "net/http"
import "html/template"

var templates = template.Must(template.ParseGlob("templates/*.tmpl"))

func handler(w http.ResponseWriter, r *http.Request) {
  type Index struct {
    // If it starts with capital letter it's public.
    // If it doesn't, it's private, so you can't use it in the templates.
    Name string // default: ""
    Title string
  }

  // data := Index { "pedro", "yooooo" }
  data := Index {
    Name: "pedro",
    Title: "yooooo gooooo!",
  }

  templates.ExecuteTemplate(w, "index", &data)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
