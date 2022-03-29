package distributor

// const (
// 	BT_HallUp   ButtonType = 0
// 	BT_HallDown            = 1
// 	BT_Cab                 = 2

func Distributor() {
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

func calculate_own_cost() {
	// Borge
	// Simple algorithm: Check own floor status + MotorDirection: cost(floor, MD)
	// More advanced:
}

func compare_delegate() {
	// Borge
	// Input:
	//	Other elevator costs
	//	Own cost
	// Output:
	//	Delegate to one elevator
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
