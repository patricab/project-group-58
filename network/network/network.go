package network

import {
	"Network-go/network/bcast"
	"Network-go/network/localip"
	"Network-go/network/peers"
	"flag"
	"fmt"
	"os"
	"time"
}


func Handler(Tx chan Msg, Rx chan Msg) {
	var id string
	var dest string
	flag.StringVar(&id, "id", "", "id of this peer")
	flag.StringVar(&dest, "dest", "", "destination id")
	flag.Parse()

	/* Default address */
	if id == "" {
		localIP, err := localip.LocalIP()
		if err != nil {
			fmt.Println(err)
			localIP = "DISCONNECTED"
		}
		id = fmt.Sprintf("%s", localIP)
	}

	// Channel for receiving updates on the id's of the peers that are
	// alive on the network
	peerUpdateCh := make(chan peers.PeerUpdate)
	// Disable/enable the transmitter after it has been started.
	peerTxEnable := make(chan bool)

	_id, _ := strconv.Atoi(id)
	_dest, _ := strconv.Atoi(dest)

	go peers.Transmitter(15648, id, peerTxEnable)
	go peers.Receiver(15648, peerUpdateCh)

	// We make channels for sending and receiving our custom data types
	tx := make(chan Msg)
	rx := make(chan Msg)

	go bcast.Transmitter(16570, tx)
	go bcast.Receiver(16570, rx)

	for{
		select {
		case m <- rx:
			if m.Dest == _id {
				Rx <- m
			}	
		}
	}
}

