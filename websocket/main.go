package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}



func main() {
	static := http.FileServer(http.Dir("./static"))
	http.Handle("/", static)

	http.HandleFunc("/echo", echo)

	log.Printf("Service started on %d \n", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error:", err.Error())
		return
	}
	log.Println("Connected...")
	defer conn.Close()
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("read message error:", err.Error())
			break
		}

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println("write message error:", err.Error())
			break
		}
	}
	log.Println("Disconnect.")
}
