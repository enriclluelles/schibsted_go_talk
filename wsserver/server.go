package wsserver

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type WsHandlers struct {
	Msg        func(msg []byte, conn *Connection, broadcast chan []byte, connections map[*Connection]bool)
	Connect    func(conn *Connection)
	Disconnect func(conn *Connection)
}

type Server struct {
	// Registered connections.
	connections map[*Connection]bool

	// Inbound messages from the connections.
	broadcast chan []byte

	// Register requests from the connections.
	register chan *Connection

	// Unregister requests from connections.
	unregister chan *Connection

	upgrader websocket.Upgrader

	Handlers WsHandlers
}

func NewServer(handlers WsHandlers) *Server {
	s := &Server{
		connections: make(map[*Connection]bool),
		broadcast:   make(chan []byte),
		register:    make(chan *Connection),
		unregister:  make(chan *Connection),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		Handlers: handlers,
	}
	go s.run()
	return s
}

func (s *Server) run() {
	for {
		select {
		case c := <-s.register:
			if s.Handlers.Connect != nil {
				s.Handlers.Connect(c)
			}
			s.connections[c] = true
		case c := <-s.unregister:
			if _, ok := s.connections[c]; ok {
				if s.Handlers.Disconnect != nil {
					s.Handlers.Disconnect(c)
				}
				delete(s.connections, c)
				close(c.send)
			}
		case m := <-s.broadcast:
			for c := range s.connections {
				select {
				case c.send <- m:
				default:
					close(c.send)
					delete(s.connections, c)
				}
			}
		}
	}
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 600 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 600 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 10000
)

type Connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

// readPump pumps messages from the websocket connection to the hub.
func (s *Server) readPump(c *Connection) {
	defer func() {
		s.unregister <- c
		c.ws.Close()
	}()
	// c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		if s.Handlers.Msg != nil {
			s.Handlers.Msg(message, c, s.broadcast, s.connections)
		}
	}
}

// write writes a message with the given message type and payload.
func (c *Connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

// writePump pumps messages from the hub to the websocket connection.
func (s *Server) writePump(c *Connection) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	c := &Connection{send: make(chan []byte, 512), ws: ws}
	s.register <- c
	go s.writePump(c)
	s.readPump(c)
}
