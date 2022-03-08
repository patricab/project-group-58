package fsm

import (
	"Driver-go/elevio"
	"fmt"
	"time"
)

/* Variables and types*/
type state int

var _dir elevio.MotorDirection
var _target elevio.ButtonEvent
var _obstruction bool
var _floor int

/* Structs and enums */
const (
	IDLE       state = 1
	MOVING     state = 2
	DOORS_OPEN state = 3
)

type current_state struct {
	idle       chan bool
	moving     chan bool
	doors_open chan bool
	active     state
}

var _current_state current_state

/* Public functions */
/////////////////////
func Handler(btn <-chan elevio.ButtonEvent, obstruction <-chan bool, stop <-chan bool, floor <-chan int) {
	for {
		/* Set global variables */
		// _obstruction = <-obstruction
		fmt.Println("In loop")
		fmt.Println("Lemme select")

		select {
		/* Stop button has been pushed */
		case s := <-stop:
			fmt.Printf("Stop: %+v\n", s)
			/* Clear button lights */
			for f := 0; f < 4; f++ {
				for b := elevio.ButtonType(0); b < 3; b++ {
					elevio.SetButtonLamp(b, f, false)
				}
			}

			_current_state.idle <- true

		case o := <-obstruction:
			fmt.Printf("Obstruction: %+v\n", o)
			set_state(DOORS_OPEN)

		case _floor = <-floor:
			fmt.Printf("Floor: %+v\n", _floor)
			elevio.SetFloorIndicator(_floor)

		case _target = <-btn:
			if _floor > _target.Floor {
				_dir = elevio.MD_Down
			} else if _floor < _target.Floor {
				_dir = elevio.MD_Up
			}
			elevio.SetMotorDirection(_dir)
			elevio.SetButtonLamp(_target.Button, _target.Floor, true)

		default:
			fmt.Println("Default case")
			fmt.Printf("Current target %+v\n", _target.Floor)

			if _floor > _target.Floor {
				set_state(MOVING)
				_dir = elevio.MD_Down
				fmt.Println("Going down")
			} else if _floor < _target.Floor {
				set_state(MOVING)
				_dir = elevio.MD_Up
				fmt.Println("Going up")
			} else if _floor == _target.Floor {
				set_state(IDLE)
				_dir = elevio.MD_Stop
				fmt.Println("Arrived at target")
			}

			elevio.SetMotorDirection(_dir)

			/* State machine */
			fmt.Printf("Current state: %+v\n", _current_state.active)
			select {
			case <-_current_state.idle:
				state_idle()
			case <-_current_state.moving:
				state_moving()
			case <-_current_state.doors_open:
				state_doors_open()
			}
		}
	}
}

/**
*  @brief Initialize FSM
*
 */
func Init() {
	/* Set initial variable values */
	// _dir = elevio.MD_Stop
	_current_state.idle = make(chan bool, 1)
	_current_state.moving = make(chan bool, 1)
	_current_state.doors_open = make(chan bool, 1)

	set_state(IDLE)
}

/* Private functions */
///////////////////////
func state_idle() {
	/* Elevator still, doors closed, waiting for floor */
	elevio.SetDoorOpenLamp(false)
	elevio.SetButtonLamp(_target.Button, _target.Floor, false)
	elevio.SetFloorIndicator(_floor)
}

func state_moving() {
	/* Elevator moving in given direction (elevio.MotorDirection) */
	elevio.SetButtonLamp(_target.Button, _target.Floor, true)
}

func state_doors_open() {
	/* Elevator still, doors open */
	elevio.SetDoorOpenLamp(true)
	time.Sleep(3 * time.Second)
	for {
		elevio.SetDoorOpenLamp(true)
		if !_obstruction {
			set_state(IDLE)
			break
		}
	}
}

/**
*  @brief Set current state of FSM
*
*  @param state Target state
 */
func set_state(target state) {
	/* Check if no state is active */
	if _current_state.active > 0 {
		/* Pop current state */
		switch _current_state.active {
		case IDLE:
			<-_current_state.idle
		case MOVING:
			<-_current_state.moving
		case DOORS_OPEN:
			<-_current_state.doors_open
		}
	}

	/* Set new state based on target */
	switch target {
	case IDLE:
		_current_state.idle <- true
	case MOVING:
		_current_state.moving <- true
	case DOORS_OPEN:
		_current_state.doors_open <- true
	}
}
