package utils

import "github.com/gorilla/websocket"

var SocketClient = make(map[string]*websocket.Conn)

func Broadcast(message string) {
	for _, conn := range SocketClient {
		conn.WriteMessage(websocket.TextMessage, []byte(message))
	}
}
func BroadcastData(data interface{}) {
	for _, conn := range SocketClient {
		conn.WriteJSON(data)
	}
}
