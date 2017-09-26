package main

import (
	"log"
	"net"

	"github.com/peterj/chat/internal/msg"
)

func main() {

	// Let's connect back and send a TCP package
	conn, err := net.Dial("tcp4", ":6000")
	if err != nil {
		log.Println("dial", err)
	}

	m := msg.MSG{
		Sender:    "Bill",
		Recipient: "Cory",
		Data:      "Hello There",
	}

	data := msg.Encode(m)

	if _, err := conn.Write(data); err != nil {
		log.Println("write", err)
	}

	if data, _, err = msg.Read(conn); err != nil {
		log.Println("read", err)
	}

	mRecv := msg.Decode(data)
	log.Println(mRecv)
}
