package main

import (
	"fmt"
	"net"

	"github.com/stvmyr/freeD"
)

var (
	addr = net.UDPAddr{
		Port: 6301,
		IP:   net.ParseIP(""),
	}
)

func main() {
	data := make([]byte, 1024)
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		_, _, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Println(err)
			continue
		}

		TrackingData, err := freeD.Decode(data)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(TrackingData)
		}

	}
}
