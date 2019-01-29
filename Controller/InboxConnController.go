package controller

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const (
	writeWait      = 10 * time.Second    // Time allowed to write a message to the peer.
	pongWait       = 60 * time.Second    // Time allowed to read the next pong message from the peer.
	pingPeriod     = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
	maxMessageSize = 48828125            // Maximum message size allowed from peer.
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// connection is an middleman between the websocket connection and the hub.
type connection struct {
	websocket *websocket.Conn // The websocket connection.
	send      chan []byte     // Buffered channel of outbound messages.
}

// readPump pumps messages from the websocket connection to the hub.
func (subscription *subscription) readPump() {
	defer func() {
		InboxHub.unregister <- *subscription
		subscription.connection.websocket.Close()
	}()
	//subscription.connection.websocket.SetReadLimit(maxMessageSize)
	subscription.connection.websocket.SetReadDeadline(time.Now().Add(pongWait))
	subscription.connection.websocket.SetPongHandler(func(string) error {
		subscription.connection.websocket.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		_, data, err := subscription.connection.websocket.ReadMessage()
		if err != nil {
			break
		}
		message := &message{
			data: data,
			room: subscription.room,
		}
		InboxHub.broadcast <- *message
	}
}

// write writes a message with the given message type and payload.
func (connection *connection) write(messageType int, payload []byte) error {
	connection.websocket.SetWriteDeadline(time.Now().Add(writeWait))
	return connection.websocket.WriteMessage(messageType, payload)
}

// writePump pumps messages from the hub to the websocket connection.
func (subscription *subscription) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		subscription.connection.websocket.Close()
	}()
	for {
		select {
		case message, ok := <-subscription.connection.send:
			if !ok {
				subscription.connection.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := subscription.connection.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := subscription.connection.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// serverWs handles websocket requests from the peer.
func ServeWs(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// get the room
	listing_id := r.URL.Query().Get("listing_id")
	if listing_id == "" {
		http.Error(w, "Invalid request", 405)
		return
	}

	connection := &connection{
		send:      make(chan []byte, 256),
		websocket: websocket,
	}

	subscription := &subscription{
		connection: connection,
		room:       listing_id,
	}

	InboxHub.register <- *subscription
	go subscription.writePump()
	subscription.readPump()
}
