package main

import (
	"fmt"
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

func reader(conn *websocket.Conn) {
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

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	reader(ws)
}

func setUpRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "simpleServer")
	})

	http.HandleFunc("/ws", serveWs)
}

func main() {
	setUpRoutes()
	http.ListenAndServe(":8080", nil)
}
