package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/TutorialEdge/realtime-chat-go-react/pkg/websocket"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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

type RoomRequest struct {
	RoomName string `json:"roomName"`
}

var (
	rooms   = make(map[string]*websocket.Pool)
	roomsMu sync.RWMutex
)

var counter int = 0

type Payload struct {
	ID       string `json:"ID"`
	RoomName string `json:"roomName"`
}

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

	r.Post("/makeRoom", func(w http.ResponseWriter, r *http.Request) {

		var room RoomRequest
		err := json.NewDecoder(r.Body).Decode(&room)

		if err != nil {
			http.Error(w, "There was an error", http.StatusBadRequest)
			return
		}

		if len(room.RoomName) == 0 {
			http.Error(w, "No room name supplied", http.StatusBadRequest)
		}

		//Start the chat room
		rooms[strconv.Itoa(counter)] = websocket.NewPool()
		go rooms[strconv.Itoa(counter)].Start()

		id := counter
		counter++

		payload := &Payload{ID: strconv.Itoa(id), RoomName: room.RoomName}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(payload)
		fmt.Printf("Room name is %s \n", room.RoomName)

	})

}

func main() {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	setUpRoutes(r)
	http.ListenAndServe(":8080", r)
}
