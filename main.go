package main

import (
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func main() {
	// Creating new server
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("Connected: ", s.ID())
		return nil
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
