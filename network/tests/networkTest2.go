package main

import (
	"Network-go/network/bcast"
	"Network-go/network/peers"
	"flag"
	"fmt"
)

// Add recieveCost struct instead of using function?

type Cmd int

const (
	cmdReqCost  Cmd = 0
	cmdCost         = 1
	cmdDelegate     = 2
)

type Msg struct {
	Id      int
	Dest    int
	Command Cmd
	Data    int // both cost and floor
}

/* Type: Initilizie
 * Desc: Inits the message system for the network
 */

// use ring for communication???

/* Type: main
 * Desc: Test for using network code
 */
func main() {

	var peerid string // id for peer
	flag.StringVar(&peerid, "id", "", "id of this peer")
	flag.Parse()

	peerUpdateCh := make(chan peers.PeerUpdate)
	PeerTxEnable := make(chan bool)

	//go bcast.Transmitter(11982, peerid, peerTxEnable)
	//go bcast.Receiver(11982, Rx)

	go peers.Transmitter(11983, peerid, PeerTxEnable)
	go peers.Receiver(11983, peerUpdateCh)

	// Gets CmdReqCost from id x with floor y
	Rx := make(chan Msg)
	Tx := make(chan Msg)

	go bcast.Transmitter(11982, Tx)
	go bcast.Receiver(11982, Rx)
	// Send cost x from distributor to id y
	var costDest int = 0
	var cost int = 2
	Msg := Msg{1, costDest, cmdCost, cost}
	Tx <- Msg
	fmt.Println("Started")

	for {
		select {
		case p := <-peerUpdateCh:
			fmt.Printf("Peer update:\n")
			fmt.Printf("  Peers:    %q\n", p.Peers)
			fmt.Printf("  New:      %q\n", p.New)
			fmt.Printf("  Lost:     %q\n", p.Lost)

		case a := <-Rx:
			fmt.Printf("Received: %#v\n", a)
		}
	}
}
