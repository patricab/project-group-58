package main

import (
	"Driver-go/elevio"
	"fsm/fsm"
)

func main() {
	elevio.Init("localhost:15657", 4)

	/* Channel declarations */
	btn := make(chan elevio.ButtonEvent)
	floor := make(chan int)

	/* Channels needed for FSM to run */
	go elevio.PollButtons(btn)
	go elevio.PollFloorSensor(floor)

	fsm.Handler(btn, floor)
}
