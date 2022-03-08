package main

import (
	"Driver-go/elevio"
	"fmt"
	"fsm/fsm"
)

// var wg = &sync.WaitGroup{}

func main() {
	elevio.Init("localhost:15657", 4)
	fsm.Init()

	/* Channel declarations */
	// btn := make(chan elevio.ButtonEvent, 3)
	btn := make(chan elevio.ButtonEvent)
	// obstr := make(chan bool, 3)
	obstr := make(chan bool)
	// floor := make(chan int, 3)
	floor := make(chan int)
	// stop := make(chan bool, 3)
	stop := make(chan bool)

	go elevio.PollObstructionSwitch(obstr)
	go elevio.PollStopButton(stop)
	go elevio.PollButtons(btn)
	go elevio.PollFloorSensor(floor)

	fmt.Println("Running Handler")
	// wg.Add(1)
	fsm.Handler(btn, obstr, stop, floor)
	// func Handler(btn <-chan elevio.ButtonEvent, obstruction <-chan bool, stop <-chan bool, floor <-chan int)
	// wg.Wait()
}
