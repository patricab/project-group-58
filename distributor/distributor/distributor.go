package distributor

import (
	"Driver-go/elevio"
	"Network-go/network/peers"
	"fmt"
	"fsm/fsm"
	"network/network"
	. "network/network"
	"strconv"
	"time"
)

/* Global variables */
var id int
var floor int

// var costArr []int
// var priorityQueue []int //[]elevio.ButtonEvent
var priorityQueue = []elevio.ButtonEvent{elevio.ButtonEvent{0, elevio.BT_Cab}}
var costArray []Msg

// var costTimeout = 500

/* Channels */
var tx = make(chan network.Msg)
var rx = make(chan network.Msg)

var _peers_chan = make(chan peers.PeerUpdate)
var _peers peers.PeerUpdate
var _btn = make(chan elevio.ButtonEvent)
var btn = make(chan elevio.ButtonEvent)

// var finished = make(chan bool)
var _floor = make(chan int)
var current_floor = make(chan int)
var current_state fsm.State

const N_ELEVATORS = 3

func Distributor(_id int) {
	/* Variables */
	// var wg sync.WaitGroup
	finished := make(chan bool)
	id = _id

	/* Initalize required modules */
	elevio.Init("localhost:15657", 4)
	go network.Handler(strconv.Itoa(_id), tx, rx, _peers_chan)
	go elevio.PollButtons(btn)
	go elevio.PollFloorSensor(_floor)
	go elevio.PollFloorSensor(current_floor)

	// wg.Add(1) // Add Handler to waitgroup
	go fsm.Handler(_btn, _floor, current_state, finished)
	// wg.Wait() // Run Handler indefinetly

	go func() {
		for {
			if len(priorityQueue) > 0 {
				fmt.Println("Sending next to FSM")
				_btn <- priorityQueue[0]
				<-finished
				fmt.Println("Reached floor")
				priorityQueue = priorityQueue[1:]
			}
		}
	}()

	go func() {
		floor = <-current_floor
	}()

	go func() {
		_peers = <-_peers_chan
	}()

	for {
		select {
		case b := <-btn:
			if b.Button == 2 { // Cab
				add_to_queue(b)
			} else { // Hall
				delegate_hall(b)
			}
			// case m := <-rx:
			// 	if m.Command == CmdDelegate {
			// 		add_to_queue(m.Data)
			// 	} else if m.Command == CmdReqCost {
			// 		cost := calculate_own_cost(m.Data)
			// 		m.Command = CmdCost // CmdCost
			// 		m.Data = cost
			// 		tx <- m
			// 	} else if m.Command == CmdCost {
			// 		costArray = append(costArray, m)
			// 	}
			// }
		}
	}
}

/*	receives cmdReqCost --> starts calculating the cost --> sends the cost back to network
 */
// func received_cmdReqCost(m Msg, dest int) {
// 	msg := Msg{id, dest, CmdCost, calculate_own_cost(m.Data)}
// 	tx <- msg
// }

func request_cost(target int) {
	// Check num of connected nodes
	numNodes := len(_peers.Peers) - 1
	//timer := time.NewTimer(time.Duration(costTimeout) * time.Millisecond)
	timeout := time.After(500 * time.Millisecond)

	// Send cost req
	tx <- Msg{id, 0, CmdReqCost, target}

	for {
		select {
		case <-timeout:
			return
		default:
			if len(costArray) == numNodes {
				return
			}
		}
	}
}

// func watchdog() {
// 	// Patric
// }

func calculate_own_cost(dest_floor int) (cost int) {

	FLOOR_TRAVEL_TIME := 2
	MOVING_PENALTY := 1
	DOOR_OPEN_TIME := 3

	switch current_state {
	case fsm.DOORS_OPEN:
		cost = int(Abs(floor-dest_floor)*FLOOR_TRAVEL_TIME + DOOR_OPEN_TIME/2)
	case fsm.IDLE:
		cost = Abs(floor-dest_floor) * FLOOR_TRAVEL_TIME
	case fsm.MOVING:
		cost = Abs(floor-dest_floor)*FLOOR_TRAVEL_TIME*FLOOR_TRAVEL_TIME + MOVING_PENALTY
	}

	return cost
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func add_to_queue(new_item elevio.ButtonEvent) {
	fmt.Println("Appending to queue")
	priorityQueue = append(priorityQueue, new_item)
}

func delegate_hall(new_item elevio.ButtonEvent) {
	// Msg
	delegate_id := 1
	delegate_dest := 1
	local_msg := Msg{delegate_id, delegate_dest, CmdCost, 0} // For comparison

	// Calculate own cost
	dest_floor := new_item.Floor
	local_msg.Data = calculate_own_cost(dest_floor)

	// Request other elevator's cost
	costArray = nil // Clear all elements
	costArray = append(costArray, local_msg)
	request_cost(dest_floor)

	// Delegate to lowest cost (Default: local)
	min_cost := local_msg.Data
	local_delegation := false

	for _, message := range costArray {
		if message.Data < min_cost {
			min_cost = message.Data
			delegate_dest = message.Dest
		} else {
			local_delegation = true
		}
	}

	if local_delegation {
		priorityQueue = append(priorityQueue, new_item)
	} else {
		msg := Msg{id, delegate_dest, CmdDelegate, dest_floor}
		tx <- msg
	}

}
