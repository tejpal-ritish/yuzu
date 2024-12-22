package pkg

import (
	"encoding/json"
	"fmt"
	"io"

	"golang.org/x/net/websocket"
)

type Chatroom interface {
}

func ReadLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection dropped: ", ws.RemoteAddr())
				break
			}
			fmt.Println("Error reading from connection: ", err)
			break
		}

		_msg := buf[:n]
		var msg Message
		err = json.Unmarshal(_msg, &msg)
		if err != nil {
			fmt.Println("Error unmarshalling message: ", err)
		}

		// logic to save the message to db tbd

		s.Broadcast(_msg)
	}
}

func Broadcast(b []byte) {
	s.mu.Lock()

	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("Write error: ", err)
				s.RemoveConn(ws)
			}
		}(ws)
	}

	s.mu.Unlock()
}
