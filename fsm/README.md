# FSM
 The finite state machine (FSM) controls the different states that an elevator is currently in allowing for specific behaviour to be broadcasted and interpreted to the rest of the system.

## States
 - ``Idle``

An elevator becomes idle when having no active orders. The idle state is briefly present when going from ``Moving`` to ``Doors open``.

 - ``Moving``

The elevator is put in a ``Moving`` state when moving between floors no matter the motor direction.

 - ``Doors open``

 When arriving at the desired floor the doors will open and put the elevator in the ``Doors open`` state. The doors are open for 3 seconds. Triggering the obstruction sensor will reset the timer.

 ## Assigning orders

 The FSM executes a order with an integer representing the desired floor.

 CLAIM: 04.04.22 - The FSM will short-circuit if more than one order is sent to it and prioritize the order sent last. Fail-safe for this must be implemented.