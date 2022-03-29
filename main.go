package main

import (
	"log"

	"golang.org/x/net/icmp"
)

func main() {
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}

	for {
		var msg []byte
		length, sourceIP, err := conn.ReadFrom(msg)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Printf("message = '%s', length = %d, source-ip = %s", string(msg), length, sourceIP)
	}
	_ = conn.Close()
}
