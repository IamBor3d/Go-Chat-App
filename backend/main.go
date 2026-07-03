package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/TutorialEdge/realtime-chat-go-react/pkg/websocket"
	"github.com/go-chi/chi/v5"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	conn, err := websocket.Upgrade(w, r)

	if err != nil {
		log.Println(w, "%+V\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

var (
	rooms   = make(map[string]*websocket.Pool)
	roomsMu sync.RWMutex
)

func setUpRoutes(r *chi.Mux) {

	r.Get("/room/{roomId}", func(w http.ResponseWriter, r *http.Request) {
		roomId := chi.URLParam(r, "roomId")
		if roomId == "" {
			http.Error(w, "Room Not Specified", http.StatusBadRequest)
			return
		}

		roomsMu.RLock()
		pool, exists := rooms[roomId]
		roomsMu.RUnlock()

		if !exists {
			http.Error(w, "Rooms Does not exist", http.StatusBadRequest)
			return
		}

		serveWs(pool, w, r)
	})

}

func main() {
	rooms["1"] = websocket.NewPool()
	rooms["2"] = websocket.NewPool()

	go rooms["1"].Start()
	go rooms["2"].Start()
	r := chi.NewRouter()
	setUpRoutes(r)
	http.ListenAndServe(":8080", r)
}
