package main

import (
	"Driver-go/elevio"
	"fmt"
)

func main() {

	// var target int
	var floor int
	var target elevio.ButtonEvent

	elevio.Init("localhost:15657", 4)

	var dir elevio.MotorDirection = elevio.MD_Stop
	elevio.SetMotorDirection(dir)

	drv_buttons := make(chan elevio.ButtonEvent)
	drv_floors := make(chan int)
	drv_obstr := make(chan bool)
	drv_stop := make(chan bool)

	go elevio.PollButtons(drv_buttons)
	go elevio.PollFloorSensor(drv_floors)
	go elevio.PollObstructionSwitch(drv_obstr)
	go elevio.PollStopButton(drv_stop)

	for {
		select {
		case target = <-drv_buttons:
			fmt.Printf("%+v\n", target)
			if floor > target.Floor {
				dir = elevio.MD_Down
			} else if floor < target.Floor {
				dir = elevio.MD_Up
			}
			elevio.SetMotorDirection(dir)
			elevio.SetButtonLamp(target.Button, target.Floor, true)

		case floor = <-drv_floors:
			fmt.Printf("%+v\n", floor)
			if floor > target.Floor {
				dir = elevio.MD_Down
			} else if floor < target.Floor {
				dir = elevio.MD_Up
			} else if floor == target.Floor {
				dir = elevio.MD_Stop
				elevio.SetButtonLamp(target.Button, target.Floor, false)
				// elevio.SetDoorOpenLamp(true)
				// time.Sleep(3)
				// elevio.SetDoorOpenLamp(false)
			}

			elevio.SetMotorDirection(dir)
			elevio.SetFloorIndicator(floor)

		case a := <-drv_obstr:
			fmt.Printf("%+v\n", a)
			// elevio.SetDoorOpenLamp(true)

		case a := <-drv_stop:
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
