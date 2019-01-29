// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package controller

type message struct {
	data []byte
	room string
}

type subscription struct {
	connection *connection
	room       string
}

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type hub struct {
	// Registered connections.
	//connections map[*connection]bool
	rooms map[string]map[*connection]bool

	// Inbound messages from the connections.
	//broadcast chan []byte
	broadcast chan message

	// Register requests from the connections.
	//register chan *connection
	register chan subscription

	// Unregister requests from connections.
	//unregister chan *connection
	unregister chan subscription
}

var InboxHub = hub{
	rooms:      make(map[string]map[*connection]bool),
	broadcast:  make(chan message),
	register:   make(chan subscription),
	unregister: make(chan subscription),
}

func (hub *hub) Run() {
	for {
		select {
		// case connection := <-hub.register:
		// 	hub.connections[connection] = true
		// case connection := <-hub.unregister:
		// 	if _, ok := hub.connections[connection]; ok {
		// 		delete(hub.connections, connection)
		// 		close(connection.send)
		// 	}
		// case message := <-hub.broadcast:
		// 	for connection := range hub.connections {
		// 		select {
		// 		case connection.send <- message:
		// 		default:
		// 			close(connection.send)
		// 			delete(hub.connections, connection)
		// 		}
		// 	}
		// }

		case subscription := <-hub.register:
			connections := hub.rooms[subscription.room]
			if connections == nil {
				connections = make(map[*connection]bool)
				hub.rooms[subscription.room] = connections
			}
			connections[subscription.connection] = true
		case subscription := <-hub.unregister:
			connections := hub.rooms[subscription.room]
			if connections != nil {
				if _, ok := connections[subscription.connection]; ok {
					delete(connections, subscription.connection)
					close(subscription.connection.send)
					if len(connections) == 0 {
						delete(hub.rooms, subscription.room)
					}
				}
			}
		case message := <-hub.broadcast:
			connections := hub.rooms[message.room]
			for connection := range connections {
				select {
				case connection.send <- message.data:
				default:
					close(connection.send)
					delete(hub.rooms[message.room], connection)
					if len(connections) == 0 {
						delete(hub.rooms, message.room)
					}
				}
			}
		}
	}
}
