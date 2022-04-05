package main

import (
	"Driver-go/elevio"
	"fmt"
	"fsm/fsm"
	"sync"
)

func main() {
	/* Variables/Channel declarations */
	var _current_state fsm.State
	var wg sync.WaitGroup
	btn := make(chan elevio.ButtonEvent)
	floor := make(chan int)
	finished := make(chan bool)

	/* Initialize elvio driver */
	elevio.Init("localhost:15657", 4)

	/* Channels needed for FSM to run */
	go elevio.PollButtons(btn)
	go elevio.PollFloorSensor(floor)

	go func() {
		for {
			select {
			case <-finished:
				fmt.Println("Reached floor")
			}
		}
	}()

	wg.Add(1) // Add Handler to waitgroup
	go fsm.Handler(btn, floor, _current_state, finished)
	wg.Wait() // Run Handler indefinetly
}
