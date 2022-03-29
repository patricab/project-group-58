package main

import (
	"Driver-go/elevio"
	"fsm/fsm"
	"sync"
)

func main() {
	/* Variables/Channel declarations */
	var wg sync.WaitGroup
	btn := make(chan elevio.ButtonEvent)
	floor := make(chan int)

	/* Initialize elvio driver */
	elevio.Init("localhost:15657", 4)

	/* Channels needed for FSM to run */
	go elevio.PollButtons(btn)
	go elevio.PollFloorSensor(floor)

	wg.Add(1) // Add Handler to waitgroup
	go fsm.Handler(btn, floor)
	wg.Wait() // Run Handler indefinetly
}
