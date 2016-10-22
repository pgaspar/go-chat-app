// Install fresh (go get github.com/pilu/fresh), cd to work dir & run `fresh`

// := does type inference
// = does not. a := 1; a = "abc" will break.
// var i float32 declares i as a float
// var i = 1 also infers

package main

import "net/http"
import "html/template"

func handler(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("templates/index.tmpl")
  t.Execute(w, nil)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
