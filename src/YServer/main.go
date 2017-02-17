package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/ws", webSocketServer)
	err := http.ListenAndServe(":9012", nil)

	if err != nil {
		log.Fatalln(err)
	}
}

func webSocketServer(w http.ResponseWriter, r *http.Request) {
	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln(err)
	}

	keepAlive(con, 5*time.Second)
}

func keepAlive(c *websocket.Conn, timeout time.Duration) {
	lastResponse := time.Now()
	c.SetPingHandler(func(msg string) error {
		log.Println("收到Ping包")
		lastResponse = time.Now()
		return nil
	})

	go func() {
		for {
			err := c.WriteMessage(websocket.PongMessage, []byte("keepalive"))
			if err != nil {
				return
			}
			log.Println("发送Pong包")
			time.Sleep(timeout / 2)
			if time.Now().Sub(lastResponse) > timeout {
				c.Close()
				return
			}
		}
	}()
}
