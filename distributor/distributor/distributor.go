package distributor

import (
	"Driver-go/elevio"
	"fmt"
	"fsm/fsm"
	"network/network"
	. "network/network"
)

/* Global variables */
// var id int

// var floor int
// var costArr []int
var priorityQueue []elevio.ButtonEvent

// var costTimeout = 500

/* Channels */
var tx = make(chan network.Msg)
var rx = make(chan network.Msg)

// var _peers = make(chan peers.PeerUpdate)
var _btn = make(chan elevio.ButtonEvent)
var btn = make(chan elevio.ButtonEvent)

// var finished = make(chan bool)
var _floor = make(chan int)
var current_state fsm.State

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
				add_to_queue(b)
			} else { // Hall
				compare_delegate(b)
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
				fmt.Println("Append cost to cost array")
				// Add cost to cost array (append(costArr, m.Data))
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

// func request_cost(target int, costReady chan bool) {
// 	// Check num of connected nodes
// 	numNodes := len(_peers.Peers) - 1

// 	// Send cost req
// 	tx <- Msg{id, 0, CmdReqCost, target}
// 	timer := time.NewTimer(time.Duration(costTimeout) * time.Millisecond)

// 	for {
// 		select {
// 		case <-timer.C:
// 			break
// 		default:
// 			if len(costArr) == numNodes {
// 				break
// 			}
// 		}
// 	}
// }

// func watchdog() {
// 	// Patric
// }

// func calculate_own_cost() (cost int) {
// 	// Borge

// 	// Switch case focuses on calculating cost intuitively based on the time it would take in seconds
// 	// Constants used
// 	//		Floor travel time (2-3 seconds?)
// 	//		Moving penalty (1)
// 	// Used variables
// 	//		Current floor
// 	//		Destination floor

// 	FLOOR_TRAVEL_TIME := 2 // sec
// 	MOVING_PENALTY := 1    // To distinguish moving elevators from idle ones, so as to not inconvenience anyone waiting for an elevator
// 	DOOR_OPEN_TIME := 3    // sec

// 	switch current_state {

// 	// case FailureMode:
// 	// 	cost = 999

// 	case fsm.DOORS_OPEN:
// 		// No obstruction
// 		cost = abs(curr_floor-dest_floor)*FLOOR_TRAVEL_TIME + DOOR_OPEN_TIME/2

// 	case fsm.IDLE:
// 		cost = abs(curr_floor-dest_floor) * FLOOR_TRAVEL_TIME

// 	case fsm.MOVING: // If its moving, it has an order
// 		// cost is distance from current floor
// 		cost = abs(curr_floor-dest_floor)*FLOOR_TRAVEL_TIME*FLOOR_TRAVEL_TIME + MOVING_PENALTY
// 	}

// 	return cost

// }

func add_to_queue(new_item elevio.ButtonEvent) {

	fmt.Println("Appending to queue")
	priorityQueue = append(priorityQueue, new_item)

	fmt.Println("Executing order")
	_btn <- priorityQueue[0]
	priorityQueue = priorityQueue[1:]
}

func compare_delegate(new_item elevio.ButtonEvent) {

	if false {
		// Add hall call to FSM?
	} else {
		fmt.Println("Appending to queue")
		priorityQueue = append(priorityQueue, new_item)

		fmt.Println("Executing order")
		_btn <- priorityQueue[0]
		priorityQueue = priorityQueue[1:]
	}
	// Compare replied cost + own cost, which is the lowest?
	// request_cost()

	// Delegate order/take order itself?
	// msg = Msg{id, dest, CmdDelegate, floor}
	// tx <- msg

	// Append order to queue (or re-evaluate queue priority?)
	// fmt.Println("Appending to queue")
	// priorityQueue = append(priorityQueue, new_item)

	// // Execute order (send to FSM)
	// fmt.Println("Executing order")
	// btn <- priorityQueue[0]

	// Cost for all elevators are added to a map
	// Example of a map after succesfully requesting cost from 3 elevators:
	//		map = {"cost_0": 1, "cost_1": 0, "cost_2": 2}
	// ... where "cost_0" always is the local elevator. This works for an arbitrary amount of elevators.

	// floor := 1          // PLACEHOLDER
	// motorDirection := 1 // PLACEHOLDER
	// elev_cost := 1      // PLACEHOLDER

	// // Create map for comparison
	// compare_map := make(map[string]int)

	// // OWN COST
	// local_cost = calculate_own_cost(floor, motorDirection)
	// compare_map["cost_0"] = local_cost

	// // GET OTHER N ELEVATOR'S COSTS
	// for i := 1; i < n_elevators+1; i++ {
	// 	// If not timed out, new cost is added to map
	// 	for !(timeout) {
	// 		elev_cost := request_cost(other_elev_port, info_about_hall_order)
	// 		compare_map["cost_"+strconv.Itoa(i)] = elev_cost
	// 	}

	// }

	// // COMPARE COST
	// min_cost := compare_map["cost_0"] // Local elevator is default
	// for elevator, _ := range compare_map {
	// 	if compare_map[elevator] < min_cost {
	// 		min_cost = compare_map
	// 	}
	// }

	// DELEGATE ORDER
	// delegate_order(min_cost) // Assuming this variable has enough info about the elevator
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
