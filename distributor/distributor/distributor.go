package distributor

import (
	"Driver-go/elevio"
	"Network-go/network/peers"
	"fmt"
	"fsm/fsm"
	"network/network"
	. "network/network"
	"time"
)

type Msg struct {
	Id      int
	Dest    int
	Command Cmd
	Data    int // both cost and floor
}

/* Global variables */
// var id int

// var floor int
// var costArr []int
var priorityQueue []int //[]elevio.ButtonEvent
var costArray []Msg

// var costTimeout = 500

/* Channels */
var tx = make(chan network.Msg)
var rx = make(chan network.Msg)

var _peers = make(chan peers.PeerUpdate)
var _btn = make(chan elevio.ButtonEvent)
var btn = make(chan elevio.ButtonEvent)

// var finished = make(chan bool)
var _floor = make(chan int)
var current_state fsm.State

const N_ELEVATORS = 3

func Distributor(_id int) {
	/* Variables */
	// var wg sync.WaitGroup
	// id = _id

	/* Initalize required modules */
	elevio.Init("localhost:15657", 4)
	// go network.Handler(strconv.Itoa(id), tx, rx, _peers)
	go elevio.PollButtons(btn)
	go elevio.PollFloorSensor(_floor)

	// wg.Add(1) // Add Handler to waitgroup
	// go fsm.Handler(_btn, _floor, current_state, finished)
	go fsm.Handler(_btn, _floor, current_state)
	// wg.Wait() // Run Handler indefinetly

	for {
		select {
		case b := <-btn:
			if b.Button == 2 { // Cab
				add_to_queue(b.Floor)
			} else { // Hall
				delegate_hall(b)
			}
		case m := <-rx:
			if m.Command == CmdDelegate {
				add_to_queue(m.Data)
			} else if m.Command == CmdReqCost {
				cost := calculate_own_cost(m.Data)
				m.Command = CmdCost // CmdCost
				m.Data = cost
				tx <- m
			} else if m.Command == CmdCost {
				costArray = append(costArray, m)
			}
		}
	}
}

/*	receives cmdReqCost --> starts calculating the cost --> sends the cost back to network
 */
// func received_cmdReqCost(dest int) {
// 	msg := Msg{id, dest, CmdCost, calculate_own_cost()}
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
				break
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
		cost = abs(curr_floor-dest_floor)*FLOOR_TRAVEL_TIME + DOOR_OPEN_TIME/2
		return cost
	case fsm.IDLE:
		cost = abs(curr_floor-dest_floor) * FLOOR_TRAVEL_TIME
		return cost
	case fsm.MOVING:
		cost = abs(curr_floor-dest_floor)*FLOOR_TRAVEL_TIME*FLOOR_TRAVEL_TIME + MOVING_PENALTY
		return cost
	}
}

func add_to_queue(new_item int) {

	fmt.Println("Appending to queue")
	priorityQueue = append(priorityQueue, new_item)

	fmt.Println("Executing order")
	_btn <- priorityQueue[0]
	priorityQueue = priorityQueue[1:]
}

func delegate_hall(new_item elevio.ButtonEvent) {

	// Msg
	local_msg := Msg{0, 0, CmdCost, 0}

	// Calculate own cost
	dest_floor := new_item.Floor
	local_msg.Data = calculate_own_cost(dest_floor)

	// Request other elevator's cost
	costArray = nil // Clear all elements
	costArray = append(costArray, local_msg)
	request_cost(dest_floor)

	// Delegate to lowest cost (Default: local)
	min_cost := local_msg.Data
	delegate_id := 99

	for _, message := range costArray {
		if message.Data < min_cost {
			min_cost = message.Data
			delegate_id = message.Id
		}
	}

	dest := 10

	msg := Msg{
		delegate_id,
		dest,
		CmdDelegate,
		dest_floor,
	}

	tx <- msg
}

// func node() {
// 	go Distributor()
// 	go Network()
// 	go FSM()
// }

// func main() {

// 	go node()
// 	go node()
// 	go node()
// }
