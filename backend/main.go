package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TutorialEdge/realtime-chat-go-react/pkg/websocket"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := websocket.Upgrade(w, r)

	if err != nil {
		log.Println(w, "%+V\n", err)
	}
	go websocket.Writer(ws)
	websocket.Reader(ws)
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
