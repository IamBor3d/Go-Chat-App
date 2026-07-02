package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Specifies parameters for upgraded from regular TCP connection to websocket
var upgrader = websocket.Upgrader{
	//The buffer size in bytes
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	//Check origin of requests
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}

	return ws, nil
}
