package main

import (
	"Network-go/network/bcast"
	"time"

	// "Network-go/network/localip"
	"Network-go/network/peers"
	"flag"
	"fmt"
	// "os"
)

/* Custom sturcts/types */
type cmd string

const (
	cmdReqCost  cmd = "0"
	cmdCost         = "1"
	cmdDelegate     = "2"
)

type Msg struct {
	id      string
	dest    string
	command cmd
	data    string // both cost and floor
}

func main() {
	var id string
	var dest string
	flag.StringVar(&id, "id", "", "id of this peer")
	flag.StringVar(&dest, "dest", "", "destination id")
	flag.Parse()

	/* Default address */
	// if _id == "" {
	// 	localIP, err := localip.LocalIP()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		localIP = "DISCONNECTED"
	// 	}
	// 	_id = fmt.Sprintf("peer-%s-%d", localIP, os.Getpid())
	// }

	// Channel for receiving updates on the id's of the peers that are
	// alive on the network
	peerUpdateCh := make(chan peers.PeerUpdate)

	// Id, _ := strconv.Atoi(id)
	// Dest, _ := strconv.Atoi(dest)

	// Disable/enable the transmitter after it has been started.
	TxEnable := make(chan bool)
	go peers.Transmitter(15648, id, TxEnable)
	go peers.Receiver(15648, peerUpdateCh)
	// TxEnable <- true

	// We make channels for sending and receiving our custom data types
	tx := make(chan Msg, 2)
	rx := make(chan Msg, 2)

	go bcast.Transmitter(16570, tx)
	go bcast.Receiver(16570, rx)

	// The example message. We just send one of these every second.
	go func() {
		fmt.Printf("id: %s\n", id)
		fmt.Printf("dest: %s\n", dest)

		msg := Msg{id, dest, cmdDelegate, "1"}
		for {
			tx <- msg
			time.Sleep(1 * time.Second)
		}
	}()

	fmt.Println("Started")
	for {
		select {
		case p := <-peerUpdateCh:
			fmt.Printf("Peer update:\n")
			fmt.Printf("  Peers:    %q\n", p.Peers)
			fmt.Printf("  New:      %q\n", p.New)
			fmt.Printf("  Lost:     %q\n", p.Lost)

		case rx_msg := <-rx:
			fmt.Println("Recieved message, checking address")
			// if rx_msg.dest == id {
			fmt.Printf("Received: %#v\n", rx_msg)
			// }
		}
	}
}
