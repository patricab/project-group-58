# Distributor

The distributor connects all the modules in the system. It is the communication hub for the system locally handling cab calls, backing up orders and delegates orders to the local FSM.

In addition to that it has two distinct roles:

1. When hall call button confined to the elevator it acts as the temporary "master" responsible for delegating a new order to the elevator with the lowest cost.

2. Transmit its own cost when inquired to do so by the elevator in charge of delegating another order.

## Cost

The cost function.

## Watchdog timer

A timer.

## Comparing & Delegating

Delegates each order to the most fitting elevator.