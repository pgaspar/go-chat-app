// Go is procedural. There are no classes, only structs (aka Interfaces)

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

// Default DB conectors will most likely work with pure SQL
// There are some ORMs available, like github.com/jinzhu/gorm

package main

import "net/http"
import "html/template"
import "github.com/gorilla/mux"
import "encoding/json"
import "github.com/gorilla/websocket"
import "github.com/pgaspar/hello_web/wslib"

var templates = template.Must(template.ParseGlob("templates/*.tmpl"))

type User struct {
  Username string `json:"user"` // Define the JSON key (instead of Username)
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

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  // Marshal goes through the whole struct tree and never includes
  // private variables
  j, _ := json.Marshal(users)

  w.Write(j)
}

func newUsersHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  username := mux.Vars(r)["username"]
  users = append(users, User {
    Username: username,
  })

  http.Redirect(w, r, "/users", 301)
}

// Websockets

var upgrader = websocket.Upgrader{}
var hub = wslib.Hub {
  Clients: make(map[*wslib.Client]bool),
  Broadcast: make(chan []byte),
  AddClient: make(chan *wslib.Client),
  RemoveClient: make(chan *wslib.Client),
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
  conn, err := upgrader.Upgrade(w, r, nil)

  if err != nil { // Be explicit when testing nil (or != nil)
    http.NotFound(w, r)
  }

  client := &wslib.Client { // We'll use this by reference, hence this
    Ws: conn,
    Send: make(chan []byte), // This defines and creates a byte stream channel
  }

  hub.AddClient <- client

  go client.Write(&hub)
  go client.Read(&hub)
}

// Main

func main() {
  go hub.Start()

  r := mux.NewRouter()
  r.HandleFunc("/", handler)
  r.HandleFunc("/users", getUsersHandler).Methods("GET")
  r.HandleFunc("/users/new/{username:[a-zA-Z]+}", newUsersHandler).Methods("GET")
  r.HandleFunc("/ws", wsHandler)

  http.ListenAndServe(":8080", r)
}
