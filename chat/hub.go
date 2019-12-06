package chat

// DefaultHub is our default hub
var DefaultHub = NewHub()

// Hub is the data structure we use to keep track of connections
type Hub struct {
	Join  chan *Conn
	Conns map[*Conn]bool
	Echo  chan string
}

// NewHub creates a new default hub.
func NewHub() *Hub {
	return &Hub{
		Join:  make(chan *Conn),
		Conns: make(map[*Conn]bool),
		Echo:  make(chan string),
	}
}

// Start starts our hub
func (hub *Hub) Start() {
	for {
		select {
		case conn := <-hub.Join:
			DefaultHub.Conns[conn] = true
		case msg := <-hub.Echo:
			for conn := range hub.Conns {
				conn.Send <- msg
			}
		}
	}
}
