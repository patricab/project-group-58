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

func received_cmdDelegate(floor chan int) {
	// Jacob
	// receives delegate order -> sends floor x to
}

func received_cmdReqCost() {
	// Jacob

}

func request_cost() {
	// Patric
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
