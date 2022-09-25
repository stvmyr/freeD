package main

import (
	"fmt"
	"net"

	"github.com/stvmyr/freeD"
)

var (
	addr = net.UDPAddr{
		Port: 6301,
		IP:   net.ParseIP("127.0.0.1"),
	}

	TrackingData = freeD.FreeD{
		Pitch: 1.257782,
		Yaw:   12.214172,
		Roll:  0.005371,
		PosX:  2532.87,
		PosY:  3274.1094,
		PosZ:  1014.3281,
		Zoom:  534,
		Focus: 1127,
	}
)

func main() {
	conn, err := net.DialUDP("udp", nil, &addr)
	if err != nil {
		fmt.Println(err)
	}

	_, err = conn.Write(freeD.Encode(TrackingData))
	if err != nil {
		fmt.Println(err)
	}
}
