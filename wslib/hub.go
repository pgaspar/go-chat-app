package wslib

// The clients will run in a different proccess than the hub.
// We communicate via channels.
type Hub struct {
  Clients map[*Client]bool
  Broadcast chan []byte
  AddClient chan *Client
  RemoveClient chan *Client
}

func (hub *Hub) Start() {
  for {
    select {
    case client := <-hub.AddClient:
      hub.Clients[client] = true
    case client := <-hub.RemoveClient:
      if _, ok := hub.Clients[client]; ok {
        delete(hub.Clients, client)
        close(client.Send)
      }
    case message := <-hub.Broadcast:
      for client := range hub.Clients {
        client.Send <- message
      }
    }
  }
}
