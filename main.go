// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var r *rand.Rand

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
		if err != nil {
			return
		}
		conn.WriteMessage(websocket.TextMessage, []byte(genCaptchaCode()))

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			fmt.Println(msgType)
			fmt.Println(websocket.BinaryMessage)
			fmt.Println(websocket.TextMessage)
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	fmt.Println("doin' your mom at :8080")
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}

func genCaptchaCode() string {
	var codes [6]byte
	for i := 0; i < 6; i++ {
		codes[i] = uint8(48 + r.Intn(10))
	}

	return string(codes[:])
}
