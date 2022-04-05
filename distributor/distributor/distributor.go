package distributor

import (
	"Driver-go/elevio"
	"fmt"
	"fsm/fsm"
	"network/network"
	. "network/network"
	"strconv"
)

/* Global variables */
var id int
var floor int
var numNodes int

// var costArr []int
// var priorityQueue []int //[]elevio.ButtonEvent
var priorityQueue = []elevio.ButtonEvent{elevio.ButtonEvent{0, elevio.BT_Cab}}
var costArray []Msg

// var costTimeout = 500

/* Channels */
var tx = make(chan network.Msg)
var rx = make(chan network.Msg)

var _btn = make(chan elevio.ButtonEvent)
var btn = make(chan elevio.ButtonEvent)

// var finished = make(chan bool)
var floor_chan = make(chan int)
var current_floor = make(chan int)
var all_costs = make(chan bool)

// var current_state fsm.State
// var current_state_chan chan fsm.State

func Distributor(_id int, port int, _numNodes int) {
	/* Variables */
	// var wg sync.WaitGroup
	finished := make(chan bool)
	id = _id
	numNodes = _numNodes
	fmt.Printf("Current ID: %d\n", id)

	/* Initalize required modules */
	host := fmt.Sprintf("localhost:%d", port)
	elevio.Init(host, 4)
	fmt.Printf("[%v] Initializing.\n", id)

	go network.Handler(strconv.Itoa(_id), tx, rx)
	go elevio.PollButtons(btn)
	go elevio.PollFloorSensor(floor_chan)
	go elevio.PollFloorSensor(current_floor)

	// wg.Add(1) // Add Handler to waitgroup
	// go fsm.Handler(_btn, floor_chan, current_state_chan, finished)
	go fsm.Handler(_btn, floor_chan, finished)
	// wg.Wait() // Run Handler indefinetly

	go func() {
		for {
			if len(priorityQueue) > 0 {
				fmt.Printf("[%v] Sending next to FSM\n", id)
				_btn <- priorityQueue[0]
				<-finished
				fmt.Printf("[%v] Reached floor\n", id)
				priorityQueue = priorityQueue[1:]
			}
		}
	}()

	go func() {
		for {
			select {
			case floor = <-current_floor:
				fmt.Printf("[%v] Floor: %d\n", id, floor)
				// case current_state = <-current_state_chan:
				// 	fmt.Printf("[%v] Current state: %d\n", id, current_state)
			}
		}
	}()

	for {
		select {
		case b := <-btn:
			if b.Button == 2 { // Cab
				add_to_queue(b)
			} else { // Hall
				go delegate_hall(b)
			}
		case m := <-rx:
			if m.Command == CmdDelegate {
				// add_to_queue(m.Data)
				fmt.Printf("[%v] Received cmdDelegate\n", id)
				add_to_queue(elevio.ButtonEvent{m.Data, elevio.BT_Cab})
			} else if m.Command == CmdReqCost {
				fmt.Printf("[%v] Received cmdReqCost\n", id)
				cost := calculate_own_cost(m.Data)
				fmt.Printf("[%v] Calculated cost: %d\n", id, cost)
				msg := Msg{id, m.Id, CmdCost, cost}
				tx <- msg
			} else if m.Command == CmdCost {
				fmt.Printf("[%v] Received cost from node\n", id)
				costArray = append(costArray, m)
				if len(costArray) == numNodes {
					fmt.Println("Received cost from all nodes")
					all_costs <- true
				}
				// fmt.Printf("[%v] Cost array: %d\n", id, costArray)
			}
		}
	}
}

/*	receives cmdReqCost --> starts calculating the cost --> sends the cost back to network
 */

// func request_cost(target int) {
// 	// Check num of connected nodes
// 	// numNodes := len(_peers.Peers) - 1
// 	//timer := time.NewTimer(time.Duration(costTimeout) * time.Millisecond)

// 	// Send cost req
// 	fmt.Printf("[%v] Sending cost request\n", id)
// 	tx <- Msg{id, 0, CmdReqCost, target}

// 	for {
// 		select {
// 		case <-time.After(3 * time.Second):
// 			fmt.Println("Timeout: request_cost")
// 			return
// 		default:
// 			if len(costArray) == numNodes {
// 				fmt.Println("Received cost from all")
// 				return
// 			}
// 		}
// 	}
// }

func calculate_own_cost(dest_floor int) (cost int) {

	FLOOR_TRAVEL_TIME := 2
	// MOVING_PENALTY := 1
	// DOOR_OPEN_TIME := 3

	fmt.Printf("[%v] Calculating cost with current floor %v and desired floor %v\n", id, floor, dest_floor)

	if len(priorityQueue) > 0 {
		cost = (Abs(floor-priorityQueue[0].Floor) + Abs(priorityQueue[0].Floor-dest_floor)) * FLOOR_TRAVEL_TIME
	} else {
		cost = (Abs(floor - dest_floor)) * FLOOR_TRAVEL_TIME
	}
	// fmt.Printf("[%v] Current state: %d\n", id, current_state)
	// switch current_state {
	// case fsm.DOORS_OPEN:
	// 	cost = int(Abs(floor-dest_floor)*FLOOR_TRAVEL_TIME + DOOR_OPEN_TIME/2)
	// case fsm.IDLE:
	// 	cost = Abs(floor-dest_floor) * FLOOR_TRAVEL_TIME
	// case fsm.MOVING:
	// 	cost = Abs(floor-dest_floor)*FLOOR_TRAVEL_TIME*FLOOR_TRAVEL_TIME + MOVING_PENALTY
	// }

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
	fmt.Printf("[%v] Appending to queue\n", id)
	priorityQueue = append(priorityQueue, new_item)
	fmt.Printf("[%v] Queue: %d\tLength: %d\n", id, priorityQueue, len(priorityQueue))
}

func delegate_hall(new_item elevio.ButtonEvent) {
	// Msg
	delegate_id := 1
	delegate_dest := 1
	local_msg := Msg{delegate_id, delegate_dest, CmdCost, 0} // For comparison

	// Calculate own cost
	dest_floor := new_item.Floor
	local_msg.Data = calculate_own_cost(dest_floor)
	fmt.Printf("[%v] Own cost is %v\n", id, local_msg.Data)

	// Request other elevator's cost
	costArray = nil // Clear all elements
	costArray = append(costArray, local_msg)
	// request_cost(dest_floor)
	tx <- Msg{id, 0, CmdReqCost, dest_floor}
	<-all_costs
	fmt.Printf("[%v] Cost array: %d\n", id, costArray)

	// Delegate to lowest cost (Default: local)
	min_cost := local_msg.Data
	local_delegation := false

	for _, message := range costArray {
		if message.Data < min_cost {
			min_cost = message.Data
			delegate_dest = message.Id
			local_delegation = false
		} else {
			local_delegation = true
		}
	}

	if local_delegation {
		fmt.Println("Taking order myself")
		priorityQueue = append(priorityQueue, new_item)
		fmt.Printf("[%v] Queue: %d\tLength: %d\n", id, priorityQueue, len(priorityQueue))
	} else {
		msg := Msg{id, delegate_dest, CmdDelegate, dest_floor}
		tx <- msg
	}

}
