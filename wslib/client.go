package wslib

import "github.com/gorilla/websocket"

type Client struct {
  Ws *websocket.Conn
  Send chan []byte // byte stream channel
}

func (c *Client) Write(hub *Hub) {
  // Always close connection whatever the path of execution
  defer func() {
    hub.RemoveClient <- c
    c.Ws.Close()
  }()

  for { // Infinite loop
    select { // select is blocking
    case msg, ok := <-c.Send:
      if !ok {
        c.Ws.WriteMessage(websocket.CloseMessage, []byte{})
        return
      }
      c.Ws.WriteMessage(websocket.TextMessage, msg)
    }
  }
}

func (c *Client) Read(hub *Hub) {
  // Always close connection whatever the path of execution
  defer func() {
    hub.RemoveClient <- c
    c.Ws.Close()
  }()

  for {
    _, msg, err := c.Ws.ReadMessage()
    if err != nil {
      return
    }

    // Broadcast message
    hub.Broadcast <- msg
  }
}
