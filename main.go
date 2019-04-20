package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {

	port := flag.Int("port", 8125, "Port that statsd metrics are being sent to")

	flag.Parse()

	addr := net.UDPAddr{
		Port: *port,
		IP: net.ParseIP("127.0.0.1"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	var buf [1024]byte
	for {
		respLen, _, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			panic(err)
		}
		fmt.Println(string(buf[:respLen]))
	}
}
