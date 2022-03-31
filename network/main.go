package main

import (
	"flag"
	"fmt"
	"network/network"
	"strconv"
	"time"
)

func main() {
	var id string
	var dest string
	flag.StringVar(&id, "id", "", "id of this peer")
	flag.StringVar(&dest, "dest", "", "destination id")
	flag.Parse()

	_id, _ := strconv.Atoi(id)
	_dest, _ := strconv.Atoi(dest)

	tx := make(chan network.Msg)
	rx := make(chan network.Msg)

	go network.Handler(id, tx, rx)

	go func() {
		fmt.Printf("id: %d\n", _id)
		fmt.Printf("dest: %d\n", _dest)

		msg := network.Msg{_id, _dest, network.CmdDelegate, 1}
		for {
			fmt.Println("Sending message")
			tx <- msg
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		select {
		case m := <-rx:
			fmt.Printf("Received: %#v\n", m)
		}
	}
}
