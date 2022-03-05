package fsm

import (
	"Driver-go/elevio"
	"fmt"
)

func Idle() {
	/* Doors closed, waiting for floor */
	// elevio.Init("localhost:15657", numFloors)
	fmt.Println("Test")
}

func Moving(dir elevio.MotorDirection) {

}

func Doors_open() {

}

func Handler() {
}
