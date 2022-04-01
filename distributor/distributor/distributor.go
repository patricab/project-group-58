package distributor

import (
	"network/network"
	"strconv"
)

var id = 0
var tx = make(chan network.Msg)
var rx = make(chan network.Msg)

// const (
// 	BT_HallUp   ButtonType = 0
// 	BT_HallDown            = 1
// 	BT_Cab                 = 2

func Distributor() {
	go network.Handler(id, tx, rx)
	for {
		select {
		case a <- drv_buttons: // Listening to PollButtons (2 is cab -> local order)
			// Check if the call is cab (local) or hall (external)
		case a: //
		}
	}
}

/*	Author: jacobkris
 *	Receives delegate order from Network --> sends floor x to state machine
 */

func received_cmdDelegate(floor int, cmdDelegate chan int) (DeleCh chan int) {

	checkDelegate := <-cmdDelegate // sets checkDelegate from channel cmdDelegate
	// set-up channel for communicating trough to the FSM
	if checkDelegate == 1 {
		DeleCh := make(chan int)
		DeleCh <- floor

	}
	return DeleCh // returns channel for sending floor x value
}

/*	Author: jacobkris
 *	receives cmdReqCost --> starts calculating the cost --> sends the cost back to network
 */
func received_cmdReqCost() (owncost int) {

	checkRq := <-rqch // gets value from command cmdReqCost channel
	if checkRq == 1 { // checks if value is one
		owncost := calculate_own_cost() // calculates own cost
	}
	return
}

/*	Author: jacobkris
 *	receives cmdReqCost --> starts calculating the cost --> sends the cost back to network
 */
func request_cost() {
	// sends out request for cost from other nodes and receives cost from other nodes
}

func watchdog() {
	// Patric
}

func calculate_own_cost(floor, MotorDirection int) {
	calculated_cost := 99 // PLACEHOLDER: high initial cost to not accidentally pass a zero as zero-value
	msg := network.Msg{id, 1, network.CmdCost, calculated_cost}
	// Borge

	// Calculate the cost for a hall call.

	// Parameters
	// cabOrders:  [bool, bool, bool, bool]
	// hallOrders: [[bool, bool, bool], [bool, bool, bool]]
	cabOrders, hallOrders := LoadBackup(backupPath) // Non-existant function for now
	desired_floor := msg.Dest

	//

	msg.cost = calculated_cost

}

func compare_delegate() {
	// Borge

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

	// DELEGATE
	delegate_order(min_cost) // Assuming this variable has enough info about the elevator
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
