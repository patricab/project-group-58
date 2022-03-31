package main

import (
	"Network-go/network/bcast"
	"Network-go/network/peers"
	"flag"
)

// Add recieveCost struct instead of using function?

type Cmd int

const (
	cmdReqCost  Cmd = 0
	cmdCost         = 1
	cmdDelegate     = 2
)

type Msg struct {
	command Cmd
	data    int // both cost and floor
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

	Rx := make(chan Msg)
	Tx := make(chan Msg)

	peerUpdateCh := make(chan peers.PeerUpdate)
	PeerTxEnable := make(chan bool)

	//go bcast.Transmitter(11982, peerid, peerTxEnable)
	//go bcast.Receiver(11982, Rx)

	go peers.Transmitter(11983, peerid, PeerTxEnable)
	go peers.Receiver(11983, peerUpdateCh)

	go bcast.Transmitter(11982, Tx)
	go bcast.Receiver(11982, Rx)

	go func() {
		msgtest := Msg{0, cmdCost, 0}
		Tx <- msgtest
	}()
}
