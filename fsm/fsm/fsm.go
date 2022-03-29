package fsm

import (
	"Driver-go/elevio"
	"fmt"
	"time"
)

/* Global variables */
var _doors_open bool = false

func Handler(button chan elevio.ButtonEvent, floor chan int) {
	/* Local variables */
	var _floor int
	var target elevio.ButtonEvent

	// elevio.Init("localhost:15657", 4)

	obstr := make(chan bool)
	stop := make(chan bool)

	var dir elevio.MotorDirection = elevio.MD_Stop
	elevio.SetMotorDirection(dir)

	go elevio.PollObstructionSwitch(obstr)
	go elevio.PollStopButton(stop)

	for {
		select {
		case target = <-button:
			fmt.Printf("%+v\n", target)
			if _floor > target.Floor {
				dir = elevio.MD_Down
			} else if _floor < target.Floor {
				dir = elevio.MD_Up
			}
			elevio.SetMotorDirection(dir)
			elevio.SetButtonLamp(target.Button, target.Floor, true)

		case _floor = <-floor:
			fmt.Printf("%+v\n", _floor)
			if _floor > target.Floor {
				dir = elevio.MD_Down
				elevio.SetMotorDirection(dir)
				elevio.SetFloorIndicator(_floor)
			} else if _floor < target.Floor {
				dir = elevio.MD_Up
				elevio.SetMotorDirection(dir)
				elevio.SetFloorIndicator(_floor)
			} else if _floor == target.Floor {
				dir = elevio.MD_Stop
				elevio.SetButtonLamp(target.Button, target.Floor, false)
				elevio.SetMotorDirection(dir)
				elevio.SetFloorIndicator(_floor)

				open_door()
			}

		case a := <-obstr:
			fmt.Printf("%+v\n", a)
			if _doors_open {
				open_door()
			}

		case a := <-stop:
			fmt.Printf("%+v\n", a)
			/* Clear button lights */
			for f := 0; f < 4; f++ {
				for b := elevio.ButtonType(0); b < 3; b++ {
					elevio.SetButtonLamp(b, f, false)
				}
			}

			/* Stop elevator! */
			elevio.SetMotorDirection(elevio.MD_Stop)
		}
	}
}

func open_door() {
	_doors_open = true
	elevio.SetDoorOpenLamp(_doors_open)
	time.Sleep(3 * time.Second)
	_doors_open = false
	elevio.SetDoorOpenLamp(_doors_open)
}
