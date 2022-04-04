package distributor

import (
	. "network/network"
	"fsm/fsm"
	"Driver-go/elevio"
	"strconv"
)

/* Global variables/channels */
var id = 0
var floor int
var tx = make(chan network.Msg)
var rx = make(chan network.Msg)
var btn = make(chan elevio.ButtonEvent)
var _floor = make(chan int)

// const (
// 	BT_HallUp   ButtonType = 0
// 	BT_HallDown            = 1
// 	BT_Cab                 = 2

func Distributor() {
	/* Variables */
	var wg sync.WaitGroup

	/* Initalize required modules */
	elevio.Init("localhost:15657", 4)
	go network.Handler(id, tx, rx)
	go elevio.PollButtons(btn)
	go elevio.PollFloorSensor(_floor)

	wg.Add(1) // Add Handler to waitgroup
	go fsm.Handler(btn, floor)
	wg.Wait() // Run Handler indefinetly

	for {
		select {
		case a <- btn: // Listening to PollButtons (2 is cab -> local order)
			// Check if the call is cab (local) or hall (external)
		case m := <-rx:
			if m.Command == CmdDelegate {
				// cmdDelegate <- m.Data
				// TODO: send data to compare/delegate
			} else if m.Command == CmdReqCost {
				received_cmdReqCost()
			}
		}
	}
}

 /*	receives cmdReqCost --> starts calculating the cost --> sends the cost back to network
 */
func received_cmdReqCost() {
	msg := Msg{id, m.Id, CmdCost, calculate_own_cost()}	
	tx <- msg
}

func request_cost() {
	// sends out request for cost from other nodes and receives cost from other nodes
	msg := Msg{id, 0, CmdReqCost, new_item}
	tx <- msg
}

func watchdog() {
	// Patric
}

func calculate_own_cost(floor, MotorDirection int) {
	// Borge

	// Switch case focuses on calculating cost intuitively based on the time it would take in seconds
	// Constants used
	//		Floor travel time (2-3 seconds?)
	//		Moving penalty (1)
	// Used variables
	//		Current floor
	//		Destination floor
	FLOOR_TRAVEL_TIME := 2 // sec
	MOVING_PENALTY := 1 // To distinguish moving elevators from idle ones, so as to not inconvenience anyone waiting for an elevator
	DOOR_OPEN_TIME := 3 // sec

	switch fsm.States {

	case FailureMode:
		cost = 999

	case DoorsOpen:
		if fsm.obstructed {
			cost = 999
		} else {
			cost = abs(curr_floor-dest_floor)*FLOOR_TRAVEL_TIME + DOOR_OPEN_TIME/2
		}

	case Idle:
		cost = abs(curr_floor-dest_floor)*FLOOR_TRAVEL_TIME

	case Moving: // If its moving, it has an order
		// cost is distance from current floor
		cost = abs(curr_floor-dest_floor)*FLOOR_TRAVEL_TIME*FLOOR_TRAVEL_TIME + MOVING_PENALTY
	}

	return cost

}

func compare_delegate(new_item int) {
	// Borge

	// Send cost request (!cmdDelegate)
	// Compare replied cost + own cost, which is the lowest?
	request_cost()

	// Delegate order/take order itself?
	msg = Msg{id, dest, CmdDelegate, floor}
	tx <- msg

	// Append order to queue (or re-evaluate queue priority?)

	// Execute order (send to FSM)
	btn <- next_order


	// COST FUNCTION
	for {
		switch
		case Idle
			cost = distance_to_floor
		case Moving
			Direction
			PriorityQueue
		case DoorOpen
			if obstructed
				high cost

		
	}

	// Cost for all elevators are added to a map
	// Example of a map after succesfully requesting cost from 3 elevators:
	//		map = {"cost_0": 1, "cost_1": 0, "cost_2": 2}
	// ... where "cost_0" always is the local elevator. This works for an arbitrary amount of elevators.

	floor := 1          // PLACEHOLDER
	motorDirection := 1 // PLACEHOLDER
	elev_cost := 1      // PLACEHOLDER

	// Create map for comparison
	compare_map := make(map[string]int)

	// OWN COST
	local_cost = calculate_own_cost(floor, motorDirection)
	compare_map["cost_0"] = local_cost

	// GET OTHER N ELEVATOR'S COSTS
	for i := 1; i < n_elevators+1; i++ {
		// If not timed out, new cost is added to map
		for !(timeout) {
			elev_cost := request_cost(other_elev_port, info_about_hall_order)
			compare_map["cost_"+strconv.Itoa(i)] = elev_cost
		}

	}

	// COMPARE COST
	min_cost := compare_map["cost_0"] // Local elevator is default
	for elevator, _ := range compare_map {
		if compare_map[elevator] < min_cost {
			min_cost = compare_map
		}
	}

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
