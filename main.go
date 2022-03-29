package main

import (
	"Driver-go/elevio"
	"fsm"
)

/* Variables and structs */

func main() {
	/* Initialize elevio module */
	elevio.Init("localhost:11982", 4)

	/* Channels */
	c_btn := make(chan elevio.ButtonEvent, 0)

	fsm.Handler()
}
