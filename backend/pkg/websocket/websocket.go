package websocket

import (
	"fmt"
	"io"
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

func Reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

func Writer(conn *websocket.Conn) {
	for {
		fmt.Print("Sending")
		messageType, r, err := conn.NextReader()

		if err != nil {
			fmt.Println(err)
			return
		}

		w, err := conn.NextWriter(messageType)

		if err != nil {
			fmt.Println(err)
			return
		}

		if _, err := io.Copy(w, r); err != nil {
			fmt.Println(err)
			return
		}

		if err := w.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}
}
