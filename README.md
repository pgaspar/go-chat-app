# Example Chat App :rocket:

Example chat app using websockets. Code and notes from a Go workshop by [@adbjesus](http://github.com/adbjesus).

## Covered Topics

Some of the topics discussed in the workshop (or that I investigated myself at the time):

#### Hello World and configuration

* Hello World
* [fresh](https://github.com/pilu/fresh)
* (Extra) Custom `runner.conf` for watching the `templates` dir
* Importing packages

#### Web Server

* Simple Hello World server
* `:=` vs. `=` and variable declaration
* Templates
* Template variables and partial templates
* Structs
* Public vs. private functions and variables
* Nested Structs as template data
* [Gorilla toolkit](http://gorillatoolkit.org)
* Using [mux](http://github.com/gorilla/mux) for routing
* In-memory user store (basically an array)
* JSON Marshalling & responding to requests with JSON
* Listing all users
* Adding users through `/users/new/<username>`
* (Extra) HTTP redirects

#### Websockets

* `websocket.Upgrader`
* Returning 404
* Constructing a Client struct for each connection
* Allow client to read and write to the websocket connection
* Add a Hub struct for managing all the connected clients
* Allow Hub to add clients, remove clients and broadcast a message to every client
* Handler action the websocket connection, instantiating a client
* Adding the client to the Hub and initializing its read and write functions
* (Extra) Moving Hub and Client code into their own files under the `wslib` package
* Instantiating the websocket in Javascript
* `onopen`, `onbeforeunload` and `onmessage`
* Listing messages on the page
* Sending messages on <enter> key press or button click
* :tada: :rocket:
