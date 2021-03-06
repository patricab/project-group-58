package network

import (
	"Network-go/network/bcast"
	"Network-go/network/localip"
	"Network-go/network/peers"
	"fmt"
	"strconv"
)

/* Custom sturcts/types */
type Cmd int

const (
	CmdReqCost  Cmd = 0 // Request cost
	CmdCost         = 1 // ACK CmdReqCost
	CmdDelegate     = 2 // Delegate order
	CmdACK          = 3 // Heartbeat
)

type Msg struct {
	Id      int
	Dest    int
	Command Cmd
	Data    int // both cost and floor
}

// func Handler(Id string, Tx chan Msg, Rx chan Msg, Peers chan peers.PeerUpdate) {
func Handler(Id string, Tx chan Msg, Rx chan Msg) {
	// var dest string
	// flag.StringVar(&id, "id", "", "id of this peer")
	// // flag.StringVar(&dest, "dest", "", "destination id")
	// flag.Parse()

	/* Default address */
	if Id == "" {
		localIP, err := localip.LocalIP()
		if err != nil {
			fmt.Println(err)
			localIP = "DISCONNECTED"
		}
		// id = fmt.Sprintf("%s", localIP)
		Id = localIP
	}

	// Channel for receiving updates on the id's of the peers that are
	// alive on the network ==> Peers

	peerUpdateCh := make(chan peers.PeerUpdate)

	// Disable/enable the transmitter after it has been started.
	peerTxEnable := make(chan bool)

	_id, _ := strconv.Atoi(Id) // Int value of ID string
	// _dest, _ := strconv.Atoi(dest)

	/* Peers Tx/Rx pair */
	go peers.Transmitter(15648, Id, peerTxEnable)
	go peers.Receiver(15648, peerUpdateCh)

	// We make channels for sending and receiving our custom data types
	// _tx := make(chan Msg)
	_rx := make(chan Msg)

	go bcast.Transmitter(16570, Tx)
	go bcast.Receiver(16570, _rx)

	for {
		select {
		case m := <-_rx:
			/* Check if recieved dest matches our ID, or special case ID (*0) */
			if (m.Dest == _id) || (m.Dest == 0 && m.Id != _id) {
				Rx <- m // Pass message on to user
			}
		}
	}
}
